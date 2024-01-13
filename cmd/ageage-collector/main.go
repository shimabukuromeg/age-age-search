package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shimabukuromeg/ageage-search/ent"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	ArticleID      string
	Title          string
	ImageURL       string
	StoreName      string
	Address        string
	SiteURL        string
	PublishedDate  string
	MunicipalityID int
}

type GeoCodeResponse struct {
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"geometry"`
	Type       string `json:"type"`
	Properties struct {
		AddressCode string `json:"addressCode"`
		Title       string `json:"title"`
	} `json:"properties"`
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type ZipcloudResponse struct {
	Message string            `json:"message"`
	Results []ZipcloudAddress `json:"results"`
	Status  int               `json:"status"`
}

type ZipcloudAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Kana1    string `json:"kana1"`
	Kana2    string `json:"kana2"`
	Kana3    string `json:"kana3"`
	PrefCode string `json:"prefcode"`
	Zipcode  string `json:"zipcode"`
}

type ExtendError struct {
	Message string
	Code    string
}

func (e *ExtendError) Error() string {
	return fmt.Sprintf("code=%s, message=%s", e.Code, e.Message)
}

func FindStoreAndAddress(siteURL string) (string, string, error) {
	// goqueryでURLからDOMオブジェクトを取得する
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return "", "", err
	}

	firstDt := doc.Find("dt").First()
	storeName := firstDt.Text()
	address := ""

	doc.Find("dt").Each(func(i int, s *goquery.Selection) {
		nfu := s.NextFilteredUntil("dd", "dt")
		if s.Text() == "住所" {
			address = nfu.Text()
		}
	})
	return storeName, address, nil
}

func FindArticles(siteURL string) ([]Article, error) {
	// goqueryでURLからDOMオブジェクトを取得する
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}
	pat := regexp.MustCompile(`.*/okitive/article/([0-9]+)/$`)
	articles := []Article{}

	doc.Find("ul li article a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 2 {
			return
		}
		title := elem.Find("p").Text()
		publishedDate := elem.Find("time").Text()
		imageURL := elem.Find("img").AttrOr("src", "")
		siteURL := elem.AttrOr("href", "")
		storeName, address, err := FindStoreAndAddress(siteURL)

		if err != nil {
			log.Fatal(err)
		}

		articles = append(articles, Article{
			ArticleID:     token[1],
			Title:         title,
			ImageURL:      imageURL,
			StoreName:     storeName,
			Address:       address,
			SiteURL:       siteURL,
			PublishedDate: publishedDate,
		})
	})

	return articles, nil
}

var dbType string
var dsn string
var target string
var isCreateSchema bool

func init() {
	flag.StringVar(&dbType, "t", "sqlite3", "Type of DB (sqlite or postgres)")
	flag.StringVar(&dsn, "d", "file:database.sqlite?_fk=1", "Database Data Source Name")
	flag.StringVar(&target, "target", "first", "target page (all or first)")
	flag.BoolVar(&isCreateSchema, "isCreateSchema", false, "execute client.Schema.Create")
	flag.Parse()

	if os.Getenv("DSN") != "" {
		dsn = os.Getenv("DSN")
		dbType = "postgres"
	}

	if dbType == "postgres" && dsn == "database.sqlite" {
		log.Fatal("When -t postgres is specified, you must specify -d with the PostgreSQL connection string. e.g. -d=postgresql://postgres@localhost:5555/ageagedb?sslmode=disable")
	}
}

func SetupDB(dbType, dsn string, isCreateSchema bool) (*ent.Client, error) {
	if dbType != "sqlite3" && dbType != "postgres" {
		return nil, fmt.Errorf("invalid dbType: %s", dbType)
	}
	client, err := ent.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}
	if isCreateSchema {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return client, nil
}

func GetMunicipalityByAddress(address string) (string, error) {
	r := regexp.MustCompile(`(沖縄県)?([^市町村]*郡)?([^市町村]*?[市町村])`)
	match := r.FindStringSubmatch(address)
	if len(match) > 3 {
		return match[3], nil // 市町村名を返す
	}
	return "", fmt.Errorf("unable to find municipality in: %s", address)
}

func GetMunicipalityByZipcode(zipcode string) (string, error) {
	baseUrl := "https://zipcloud.ibsnet.co.jp/api/search?zipcode="
	resp, err := http.Get(baseUrl + zipcode)
	if err != nil {
		return "", fmt.Errorf("failed getting response from zipcloud.ibsnet.co.jp: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("fail ioutil.ReadAll : %v", err)
	}

	// レスポンスを構造体にマッピング
	var response ZipcloudResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("fail Unmarshal : %v", err)
	}

	if response.Results == nil {
		return "", &ExtendError{"resultsの結果がnullです", "zipcode-not-found"}
	}

	if len(response.Results[0].Address2) == 0 {
		return "", fmt.Errorf("no result found for result: %s", response.Results[0])
	}

	return response.Results[0].Address2, nil
}

func GetPostalAndAddress(fullAddress string) (string, string, error) {
	r := regexp.MustCompile(`〒([0-9]{3})-([0-9]{4})\s?(.*)`)
	match := r.FindStringSubmatch(fullAddress)
	if len(match) > 3 {
		postalCode := match[1] + match[2]      // Postal code
		address := strings.TrimSpace(match[3]) // Address
		return postalCode, address, nil
	}
	return "", "", fmt.Errorf("unable to find postal code and address in: %s", fullAddress)
}

func CreateMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Municipality, error) {
	zipCode, address, err := GetZipcodeAndAddress(article.Address)

	fmt.Printf("zipCode: %s\n", zipCode)
	if err != nil {
		fmt.Println(err)
	}

	name, err := GetMunicipalityByAddress(address)
	if err != nil {
		n, err := GetMunicipalityByZipcode(zipCode)
		if err != nil {
			return nil, fmt.Errorf("failed exec GetMunicipalityByZipcode: %w", err)
		}
		name = n
	}
	id, err := client.Municipality.
		Create().
		SetName(name).
		SetZipcode(zipCode).
		OnConflictColumns("name").
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating municipality: %w", err)
	}

	createdMunicipality, err := client.Municipality.
		Query().
		Where(municipality.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying the municipality: %w", err)
	}

	return createdMunicipality, nil
}

func GetZipcodeAndAddress(fullAddress string) (string, string, error) {
	r := regexp.MustCompile(`〒([0-9]{3})-([0-9]{4})\s?(.*)`)
	match := r.FindStringSubmatch(fullAddress)
	if len(match) > 3 {
		zipCode := match[1] + match[2]         // Postal code
		address := strings.TrimSpace(match[3]) // Address
		return zipCode, address, nil
	}
	return "", "", fmt.Errorf("unable to find postal code and address in: %s", fullAddress)
}

func CreateMeshiAndMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error) {

	zipCode, address, err := GetZipcodeAndAddress(article.Address)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(zipCode, address)
	}
	location, err := GetLatLng(address)
	if err != nil {
		return nil, fmt.Errorf("fail get latlng: %w", err)
	}
	fmt.Println(location)

	municipality, err := CreateMunicipality(context.Background(), client, article)
	if err != nil {
		log.Println("fail crate municipality: %w", err)
	}

	layout := "2006/01/02"
	parsedTime, err := time.Parse(layout, article.PublishedDate)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	id, err := client.Meshi.
		Create().
		SetArticleID(article.ArticleID).
		SetTitle(article.Title).
		SetImageURL(article.ImageURL).
		SetStoreName(article.StoreName).
		SetAddress(article.Address).
		SetSiteURL(article.SiteURL).
		SetPublishedDate(parsedTime).
		SetLatitude(location.Latitude).
		SetLongitude(location.Longitude).
		OnConflictColumns("article_id").
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating meshi: %w", err)
	}

	createdMeshi, err := client.Meshi.
		Query().
		Where(meshi.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying the meshi: %w", err)
	}

	// Link the meshi to the municipality
	err = client.Meshi.
		UpdateOneID(createdMeshi.ID).
		SetMunicipalityID(municipality.ID).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed linking the meshi to the municipality: %w", err)
	}

	return createdMeshi, nil
}

func GetLatLng(address string) (Location, error) {
	baseUrl := "https://msearch.gsi.go.jp/address-search/AddressSearch"

	params := url.Values{}
	params.Add("q", address)

	resp, err := http.Get(baseUrl + "?" + params.Encode())
	if err != nil {
		return Location{}, fmt.Errorf("failed AddressSearch request : %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Location{}, fmt.Errorf("fail ioutil.ReadAll : %v", err)
	}

	var data []GeoCodeResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return Location{}, fmt.Errorf("fail Unmarshal : %v", err)
	}

	if len(data) == 0 {
		return Location{}, fmt.Errorf("no result found for address: %s", address)
	}

	latLng := Location{
		Latitude:  data[0].Geometry.Coordinates[1],
		Longitude: data[0].Geometry.Coordinates[0],
	}
	return latLng, nil
}

func main() {
	flag.Parse()
	client, err := SetupDB(dbType, dsn, isCreateSchema)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	cxt := context.Background()
	baseURL := "https://www.otv.co.jp/okitive/collaborator/ageage/page/%d"
	page := 1
	for {
		listURL := fmt.Sprintf(baseURL, page)
		articles, err := FindArticles(listURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found %d articles", len(articles))
		if len(articles) == 0 {
			break
		}
		for _, article := range articles {
			fmt.Println(article)
			_, err := client.Meshi.
				Query().
				Where(meshi.ArticleIDEQ(article.ArticleID)).
				Only(cxt)
			if err != nil {
				// 存在しなかったら追加
				if ent.IsNotFound(err) {
					fmt.Println("No Meshi found with this articleID.")
					// Create a new Meshi and link it to the Municipality.
					_, err := CreateMeshiAndMunicipality(cxt, client, &article)
					if err != nil {
						log.Println("fail crate meshi: %w", err)
						continue
					}
				}
			}

			time.Sleep(time.Second * 1)
		}
		if target == "first" {
			break
		}
		page++
	}
	fmt.Println("done")
}

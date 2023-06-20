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
	"github.com/shimabukuromeg/ageage-search/ent/migrate"
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

func init() {
	flag.StringVar(&dbType, "t", "sqlite3", "Type of DB (sqlite or postgres)")
	flag.StringVar(&dsn, "d", "file:database.sqlite?_fk=1", "Database Data Source Name")
	flag.StringVar(&target, "target", "first", "target page (all or first)")
	flag.Parse()

	if os.Getenv("DSN") != "" {
		dsn = os.Getenv("DSN")
		dbType = "postgres"
	}

	if dbType == "postgres" && dsn == "database.sqlite" {
		log.Fatal("When -t postgres is specified, you must specify -d with the PostgreSQL connection string. e.g. -d=postgresql://postgres@localhost:5555/ageagedb?sslmode=disable")
	}
}

func SetupDB(dbType, dsn string) (*ent.Client, error) {
	client, err := ent.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}
	if dbType != "sqlite3" && dbType != "postgres" {
		return nil, fmt.Errorf("invalid dbType: %s", dbType)
	}
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}

func GetMunicipality(address string) (string, error) {
	r := regexp.MustCompile(`沖縄県([^市町村]*?[市町村])`)
	match := r.FindStringSubmatch(address)
	if len(match) > 1 {
		return match[1], nil // 市町村名を返す
	}
	return "", fmt.Errorf("unable to find municipality in: %s", address)
}

func CreateMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Municipality, error) {
	name, err := GetMunicipality(article.Address)
	if err != nil {
		return nil, fmt.Errorf("failed getting municipalityName: %w", err)
	}
	id, err := client.Municipality.
		Create().
		SetName(name).
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

func CreateMeshiAndMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error) {

	// article.Address には〒901-0242　沖縄県豊見城市高安576-1 が入っているので住所だけし抽出する
	newAddress := strings.Replace(article.Address, "　", " ", -1)
	address := strings.Split(newAddress, " ")[1]
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
	client, err := SetupDB(dbType, dsn)
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

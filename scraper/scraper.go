package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shimabukuromeg/ageage-search/ent"
)

// ExtendError はエラー情報を拡張するための構造体
type ExtendError struct {
	Message string
	Code    string
}

func (e *ExtendError) Error() string {
	return fmt.Sprintf("code=%s, message=%s", e.Code, e.Message)
}

// Location は位置情報を表す構造体
type Location struct {
	Latitude  float64
	Longitude float64
}

// GeoCodeResponse は国土地理院APIからのレスポンスを表す構造体
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

// ZipcloudResponse はZipcloud APIからのレスポンスを表す構造体
type ZipcloudResponse struct {
	Message string            `json:"message"`
	Results []ZipcloudAddress `json:"results"`
	Status  int               `json:"status"`
}

// ZipcloudAddress はZipcloud APIのアドレス情報を表す構造体
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

// FindStoreAndAddress は記事URLから店名と住所を取得します
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

// FindArticles は記事一覧から記事情報を取得します
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
			log.Printf("Failed to find store and address: %v", err)
			return
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

// GetMunicipalityByAddress は住所から市町村名を抽出します
func GetMunicipalityByAddress(address string) (string, error) {
	r := regexp.MustCompile(`(沖縄県)?([^市町村]*郡)?([^市町村]*?[市町村])`)
	match := r.FindStringSubmatch(address)
	if len(match) > 3 {
		return match[3], nil // 市町村名を返す
	}
	return "", fmt.Errorf("unable to find municipality in: %s", address)
}

// GetMunicipalityByZipcode は郵便番号から市町村名を取得します
func GetMunicipalityByZipcode(zipcode string) (string, error) {
	baseUrl := "https://zipcloud.ibsnet.co.jp/api/search?zipcode="
	resp, err := http.Get(baseUrl + zipcode)
	if err != nil {
		return "", fmt.Errorf("failed getting response from zipcloud.ibsnet.co.jp: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("fail ReadAll : %v", err)
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
		return "", fmt.Errorf("no result found for result: %+v", response.Results[0])
	}

	return response.Results[0].Address2, nil
}

// GetZipcodeAndAddress は住所から郵便番号と住所を抽出します
func GetZipcodeAndAddress(fullAddress string) (string, string, error) {
	// 〒マークがあるパターンと、ないパターンの両方に対応
	r := regexp.MustCompile(`(?:〒)?([0-9]{3})-([0-9]{4})\s?(.*)`)
	match := r.FindStringSubmatch(fullAddress)
	if len(match) > 3 {
		zipCode := match[1] + match[2]         // 郵便番号
		address := strings.TrimSpace(match[3]) // 住所
		return zipCode, address, nil
	}
	return "", "", fmt.Errorf("unable to find postal code and address in: %s", fullAddress)
}

// CreateMunicipality は市町村データを作成します
func CreateMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Municipality, error) {
	zipCode, address, err := GetZipcodeAndAddress(article.Address)

	fmt.Printf("zipCode: %s\n", zipCode)
	if err != nil {
		fmt.Println(err)
		// 郵便番号が取得できなくても、住所から市町村を抽出する
		name, err := GetMunicipalityByAddress(article.Address)
		if err != nil {
			return nil, fmt.Errorf("failed to extract municipality from address: %w", err)
		}
		
		// 市町村が取得できたらそれを使用
		id, err := client.Municipality.
			Create().
			SetName(name).
			SetZipcode("unknown").  // 郵便番号不明の場合
			OnConflict().
			UpdateNewValues().
			ID(ctx)
		
		if err != nil {
			return nil, fmt.Errorf("failed creating municipality: %w", err)
		}
		
		// 作成または更新されたMunicipalityを取得
		municipality, err := client.Municipality.Get(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed getting municipality after create: %w", err)
		}
		
		return municipality, nil
	}

	name, err := GetMunicipalityByAddress(address)
	if err != nil {
		n, err := GetMunicipalityByZipcode(zipCode)
		if err != nil {
			// Zipcloudでもエラーになったらaddressをそのまま使用して市町村名を抽出する
			name, err := GetMunicipalityByAddress(article.Address)
			if err != nil {
				// 最終手段として"沖縄県"をデフォルト値として使う
				name = "沖縄県"
				log.Printf("Using default municipality 沖縄県 for article %s", article.ArticleID)
			}
			
			id, err := client.Municipality.
				Create().
				SetName(name).
				SetZipcode(zipCode).
				OnConflict().
				UpdateNewValues().
				ID(ctx)
				
			if err != nil {
				return nil, fmt.Errorf("failed creating municipality: %w", err)
			}
			
			// 作成または更新されたMunicipalityを取得
			municipality, err := client.Municipality.Get(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("failed getting municipality after create: %w", err)
			}
			
			return municipality, nil
		}
		name = n
	}
	id, err := client.Municipality.
		Create().
		SetName(name).
		SetZipcode(zipCode).
		OnConflict().
		UpdateNewValues().
		ID(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating municipality: %w", err)
	}

	// 作成または更新されたMunicipalityを取得
	municipality, err := client.Municipality.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting municipality after create: %w", err)
	}

	return municipality, nil
}

// GetLatLng は住所から緯度経度を取得します
func GetLatLng(address string) (Location, error) {
	latLng := Location{}
	baseUrl := "https://msearch.gsi.go.jp/address-search/AddressSearch?q="
	
	// URLエンコードされた住所をリクエスト
	resp, err := http.Get(baseUrl + address)
	if err != nil {
		return latLng, fmt.Errorf("failed getting response from gsi.go.jp: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return latLng, fmt.Errorf("fail ReadAll : %v", err)
	}

	// レスポンスJSONを解析
	var geoResults []GeoCodeResponse
	err = json.Unmarshal(body, &geoResults)
	if err != nil {
		return latLng, fmt.Errorf("fail Unmarshal : %v", err)
	}

	if len(geoResults) > 0 {
		// 座標系が逆（GeoJSONは[経度, 緯度]の順）
		latLng.Longitude = geoResults[0].Geometry.Coordinates[0]
		latLng.Latitude = geoResults[0].Geometry.Coordinates[1]
	}
	return latLng, nil
}

// CreateMeshiAndMunicipality は飲食店データと市町村データを作成します
func CreateMeshiAndMunicipality(ctx context.Context, client *ent.Client, article *Article) (*ent.Meshi, error) {
	// 1. Municipalityを作成
	municipality, err := CreateMunicipality(ctx, client, article)
	if err != nil {
		return nil, fmt.Errorf("failed creating municipality: %w", err)
	}

	// 2. 日付文字列をパース
	layout := "2006.01.02"
	publishedDate, err := time.Parse(layout, article.PublishedDate)
	if err != nil {
		publishedDate = time.Now() // エラー時はデフォルト値
	}

	// 3. 緯度・経度を取得
	_, address, err := GetZipcodeAndAddress(article.Address)
	if err != nil {
		address = article.Address // 郵便番号フォーマットでない場合はそのまま使用
	}
	loc, err := GetLatLng(address)
	if err != nil {
		// エラーがあってもデフォルト値で続行
		loc = Location{Latitude: 0, Longitude: 0}
	}

	// 4. Meshiを作成し、Municipalityと関連付け
	meshi, err := client.Meshi.
		Create().
		SetArticleID(article.ArticleID).
		SetTitle(article.Title).
		SetImageURL(article.ImageURL).
		SetStoreName(article.StoreName).
		SetAddress(article.Address).
		SetSiteURL(article.SiteURL).
		SetPublishedDate(publishedDate).
		SetLatitude(loc.Latitude).
		SetLongitude(loc.Longitude).
		SetMunicipality(municipality).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating meshi: %w", err)
	}

	return meshi, nil
}

// SetupDB はデータベース接続を設定します
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

// Runner はスクレイピング処理を実行します
func Runner(client *ent.Client, target string, numWorkers int, rps float64, limit int) error {
	log.Printf("Starting scraper with %d workers, rate limit: %.1f requests/sec", numWorkers, rps)
	
	// リミットとターゲットのログ表示
	if limit > 0 {
		log.Printf("Limit set to %d articles", limit)
	}
	
	if target == "single" {
		log.Printf("Target mode: single page (only first page)")
	} else if target == "all" {
		log.Printf("Target mode: all pages")
	} else {
		log.Printf("Unknown target mode: %s, defaulting to single page", target)
		target = "single"
	}
	
	// ワーカープールの設定
	pool := NewWorkerPool(numWorkers, client, rps, CreateMeshiAndMunicipality)
	pool.Start()
	
	// 結果処理用変数
	successCount := 0
	errorCount := 0
	totalProcessed := 0
	
	// 記事一覧の取得
	baseURL := "https://www.otv.co.jp/okitive/collaborator/ageage/page/%d"
	page := 1
	totalQueued := 0
	
	// 処理完了したページ数
	pagesProcessed := 0
	
	// 結果チャネルを処理するgoroutine
	resultChan := make(chan *ProcessResult, 100) // 結果を一時的に保存するチャネル
	
	go func() {
		for result := range resultChan {
			if result.Error != nil {
				errorCount++
				log.Printf("Error processing article %s: %v", 
						  result.Article.ArticleID, result.Error)
			} else {
				successCount++
				log.Printf("Successfully processed article %s", result.Article.ArticleID)
			}
			totalProcessed++
			log.Printf("Progress: %d successful, %d errors of %d processed so far", 
				successCount, errorCount, totalProcessed)
		}
	}()
	
	for {
		// 制限に達したかチェック
		if limit > 0 && totalQueued >= limit {
			log.Printf("Reached limit of %d articles, stopping article collection", limit)
			break
		}
		
		listURL := fmt.Sprintf(baseURL, page)
		log.Printf("Fetching articles from %s", listURL)
		
		// ページから記事を取得
		pageArticles, err := FindArticles(listURL)
		if err != nil {
			log.Printf("Error finding articles: %v", err)
			break
		}
		
		log.Printf("Found %d articles on page %d", len(pageArticles), page)
		
		// 記事が見つからなければ終了
		if len(pageArticles) == 0 {
			break
		}
		
		// このページの記事を処理
		articlesToProcess := pageArticles
		
		// 制限に達する場合は部分的に処理
		if limit > 0 {
			remaining := limit - totalQueued
			if remaining <= 0 {
				break
			}
			if len(articlesToProcess) > remaining {
				articlesToProcess = articlesToProcess[:remaining]
			}
		}
		
		// このページの処理を待機するための変数
		pageComplete := make(chan bool)
		pageArticlesCount := len(articlesToProcess)
		pageProcessed := 0
		
		// このページの結果処理用goroutine
		go func() {
			// このページの記事がすべて処理されるまで待機
			for result := range pool.ResultChan {
				// 結果を転送
				resultChan <- result
				
				pageProcessed++
				if pageProcessed >= pageArticlesCount {
					// このページのすべての記事が処理された
					close(pageComplete)
					break
				}
			}
		}()
		
		// 各記事をキューに追加
		for i := range articlesToProcess {
			// 値渡しでコピーを作成して使用
			article := articlesToProcess[i]
			
			// ジョブキューに送信
			pool.JobQueue <- &article
			log.Printf("Queued article: %s - %s", article.ArticleID, article.Title)
			totalQueued++
		}
		
		// このページの処理が完了するのを待つ
		<-pageComplete
		pagesProcessed++
		log.Printf("Completed processing all articles from page %d", page)
		
		// 単一ページモードの場合は最初のページのみ処理
		if target == "single" {
			log.Printf("Single page mode: stopping after first page")
			break
		}
		
		page++
		
		// ページ間の移動にも適切な間隔を設ける
		log.Printf("Waiting 2 seconds before fetching next page...")
		time.Sleep(time.Second * 2)
	}
	
	// すべての記事キューイングが完了
	log.Printf("All articles queued and processed (%d total across %d pages)", 
		totalProcessed, pagesProcessed)
	
	// チャネルを閉じる
	close(resultChan)
	close(pool.JobQueue)
	close(pool.ResultChan)
	
	log.Printf("All jobs completed. Processed %d articles: %d successful, %d errors", 
		totalProcessed, successCount, errorCount)
	
	return nil
} 
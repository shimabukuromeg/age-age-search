package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Meshi struct {
	ArticleID string
	Title     string
	ImageURL  string
	StoreName string
	Address   string
	SiteURL   string
	MunicipalityID int
}

type Municipality struct {
	ID   int
	Name string
}

func findStoreAndAddress(siteURL string) (string, string, error) {
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

func findMeshis(siteURL string) ([]Meshi, error) {
	// goqueryでURLからDOMオブジェクトを取得する
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}
	pat := regexp.MustCompile(`.*/okitive/article/([0-9]+)/$`)
	meshis := []Meshi{}

	doc.Find("ul li article a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 2 {
			return
		}
		title := elem.Find("p").Text()
		imageURL := elem.Find("img").AttrOr("src", "")
		siteURL := elem.AttrOr("href", "")
		storeName, address, err := findStoreAndAddress(siteURL)

		if err != nil {
			log.Fatal(err)
		}

		meshis = append(meshis, Meshi{
			ArticleID: token[1],
			Title:     title,
			ImageURL:  imageURL,
			StoreName: storeName,
			Address:   address,
			SiteURL:   siteURL,
		})
	})

	return meshis, nil
}

var dbType string
var dsn string
var target string

func init() {
	flag.StringVar(&dbType, "t", "sqlite3", "Type of DB (sqlite or postgres)")
	flag.StringVar(&dsn, "d", "database.sqlite", "Database Data Source Name")
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

func setupDB(dbType, dsn string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	var queries []string

	switch dbType {
	case "sqlite3":
		db, err = sql.Open("sqlite3", dsn)
		if err != nil {
			return nil, err
		}

		queries = []string{
			`CREATE TABLE IF NOT EXISTS municipalities(
				id INTEGER PRIMARY KEY,
				name TEXT NOT NULL UNIQUE
			)`,
			`CREATE TABLE IF NOT EXISTS meshis(
				id INTEGER PRIMARY KEY,
				article_id TEXT NOT NULL UNIQUE,
				title TEXT,
				image_url TEXT,
				store_name TEXT,
				address TEXT,
				site_url TEXT,
				municipality_id INTEGER,
				FOREIGN KEY(municipality_id) REFERENCES municipalities(id)
			)`,
		}
	case "postgres":
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return nil, err
		}

		queries = []string{
			`CREATE TABLE IF NOT EXISTS municipalities(
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL UNIQUE
			)`,
			`CREATE TABLE IF NOT EXISTS meshis(
				id SERIAL PRIMARY KEY,
				article_id TEXT NOT NULL UNIQUE,
				title TEXT,
				image_url TEXT,
				store_name TEXT,
				address TEXT,
				site_url TEXT,
				municipality_id INTEGER,
				FOREIGN KEY(municipality_id) REFERENCES municipalities(id)
			)`,
		}
	default:
		return nil, fmt.Errorf("Unknown database type: %s", dbType)
	}

	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func getMunicipality(address string) (string, error) {
	r := regexp.MustCompile(`沖縄県([^市町村]*?[市町村])`)
	match := r.FindStringSubmatch(address)
	if len(match) > 1 {
		return match[1], nil // 市町村名を返す
	}
	return "", fmt.Errorf("unable to find municipality in: %s", address)
}

func addMeshi(db *sql.DB, meshi *Meshi) error {
	municipality, err := getMunicipality(meshi.Address)
	if err != nil {
		return err
	}
	var municipalityID int
	switch dbType {
	case "sqlite3":
		err = db.QueryRow("SELECT id FROM municipalities WHERE name = ?", municipality).Scan(&municipalityID)
		if err == sql.ErrNoRows {
			// If the municipality doesn't exist, insert it and get its ID
			res, err := db.Exec("INSERT INTO municipalities(name) VALUES(?)", municipality)
			if err != nil {
				return err
			}
			lastID, err := res.LastInsertId()
			if err != nil {
				return err
			}
			municipalityID = int(lastID)
		} else if err != nil {
			// If another error occurred, return it
			return err
		}

		_, err = db.Exec(`
	        REPLACE INTO meshis(article_id, title, image_url, store_name, address, site_url, municipality_id) VALUES(?, ?, ?, ?, ?, ?, ?)
	    `,
			meshi.ArticleID,
			meshi.Title,
			meshi.ImageURL,
			meshi.StoreName,
			meshi.Address,
			meshi.SiteURL,
			municipalityID,
		)
		if err != nil {
			return err
		}

	case "postgres":
		err = db.QueryRow("SELECT id FROM municipalities WHERE name = $1", municipality).Scan(&municipalityID)
		if err == sql.ErrNoRows {
			// If the municipality doesn't exist, insert it and get its ID
			err = db.QueryRow("INSERT INTO municipalities(name) VALUES($1) RETURNING id", municipality).Scan(&municipalityID)
			if err != nil {
				return err
			}
		} else if err != nil {
			// If another error occurred, return it
			return err
		}

		// PostgreSQL does not support the REPLACE statement, so we use an INSERT statement with the ON CONFLICT DO UPDATE clause.
		_, err = db.Exec(`
	        INSERT INTO meshis(article_id, title, image_url, store_name, address, site_url, municipality_id) VALUES($1, $2, $3, $4, $5, $6, $7)
	        ON CONFLICT (article_id) DO UPDATE SET title = excluded.title, image_url = excluded.image_url, store_name = excluded.store_name, address = excluded.address, site_url = excluded.site_url, municipality_id = excluded.municipality_id
	    `,
			meshi.ArticleID,
			meshi.Title,
			meshi.ImageURL,
			meshi.StoreName,
			meshi.Address,
			meshi.SiteURL,
			municipalityID,
		)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unknown database type: %s", dbType)
	}

	return nil
}

func main() {
	flag.Parse()

	var db *sql.DB
	var err error

	if dbType == "sqlite3" || dbType == "postgres" {
		db, err = setupDB(dbType, dsn)
	} else {
		log.Fatalf("Unsupported DB type '%s'. Only 'sqlite' and 'postgres' are supported.", dbType)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	baseURL := "https://www.otv.co.jp/okitive/collaborator/ageage/page/%d"
	page := 1
	for {
		listURL := fmt.Sprintf(baseURL, page)
		meshis, err := findMeshis(listURL)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found %d meshis", len(meshis))
		if len(meshis) == 0 {
			break
		}
		for _, Meshi := range meshis {
			fmt.Println(Meshi.ArticleID)
			fmt.Println(Meshi.Title)
			fmt.Println(Meshi.ImageURL)
			fmt.Println(Meshi.StoreName)
			fmt.Println(Meshi.Address)
			fmt.Println(Meshi.SiteURL)

			err = addMeshi(db, &Meshi)
			if err != nil {
				log.Println(err)
				continue
			}

			time.Sleep(time.Second * 1)
		}
		if (target == "first") {
			break
		}
		page++
	}
	fmt.Println("done")
}

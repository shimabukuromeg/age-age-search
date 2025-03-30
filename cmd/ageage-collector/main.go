package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shimabukuromeg/ageage-search/scraper"
)

var dbType string
var dsn string
var target string
var isCreateSchema bool
var numWorkers int
var rps float64

func init() {
	flag.StringVar(&dbType, "t", "sqlite3", "Type of DB (sqlite or postgres)")
	flag.StringVar(&dsn, "d", "file:database.sqlite?_fk=1", "Database Data Source Name")
	flag.StringVar(&target, "target", "first", "target page (all or first)")
	flag.BoolVar(&isCreateSchema, "isCreateSchema", false, "execute client.Schema.Create")
	flag.IntVar(&numWorkers, "workers", 8, "Number of worker goroutines")
	flag.Float64Var(&rps, "rps", 2.0, "Requests per second (rate limit)")
	flag.Parse()

	// 環境変数からDSNを取得
	if os.Getenv("DSN") != "" {
		dsn = os.Getenv("DSN")
		dbType = "postgres"
	}

	// Postgresの場合、DSNが必要
	if dbType == "postgres" && dsn == "file:database.sqlite?_fk=1" {
		log.Fatal("When -t postgres is specified, you must specify -d with the PostgreSQL connection string. e.g. -d=postgresql://postgres@localhost:5555/ageagedb?sslmode=disable")
	}
}

func main() {
	// データベース接続を設定
	client, err := scraper.SetupDB(dbType, dsn, isCreateSchema)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 並列スクレイピングを実行
	err = scraper.Runner(client, target, numWorkers, rps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}

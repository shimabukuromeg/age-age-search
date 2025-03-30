package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ageagesearch "github.com/shimabukuromeg/ageage-search"
	"github.com/shimabukuromeg/ageage-search/ent"
	"github.com/shimabukuromeg/ageage-search/ent/migrate"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

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

func main() {
	if dbType != "sqlite3" && dbType != "postgres" {
		log.Fatal("opening ent client", dbType)
	}
	client, err := ent.Open(dbType, dsn)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	// Create ent.Client and run the schema migration.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(ageagesearch.NewSchema(client))
	http.Handle("/",
		playground.Handler("Ageagesearch", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

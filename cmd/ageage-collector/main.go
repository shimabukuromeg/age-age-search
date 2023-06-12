package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	ArticleID string
	Title     string
	ImageURL  string
	StoreName string
	Address   string
	SiteURL   string
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

func findEntries(siteURL string) ([]Entry, error) {
	// goqueryでURLからDOMオブジェクトを取得する
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}
	pat := regexp.MustCompile(`.*/okitive/article/([0-9]+)/$`)
	entries := []Entry{}

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

		entries = append(entries, Entry{
			ArticleID: token[1],
			Title:     title,
			ImageURL:  imageURL,
			StoreName: storeName,
			Address:   address,
			SiteURL:   siteURL,
		})
	})

	return entries, nil
}

func main() {
	listURL := "https://www.otv.co.jp/okitive/collaborator/ageage/page/1"
	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("found %d entries", len(entries))
	for _, entry := range entries {
		fmt.Println(entry.ArticleID)
		fmt.Println(entry.Title)
		fmt.Println(entry.ImageURL)
		fmt.Println(entry.StoreName)
		fmt.Println(entry.Address)
		fmt.Println(entry.SiteURL)

		time.Sleep(time.Second * 1)
	}
}
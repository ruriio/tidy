package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	. "tidy/model"
	. "tidy/sites"
)

func Scrape(site Site) Meta {
	var meta = Meta{}
	log.Printf("url: %s", site.Url)
	resp, err := get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("stats code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	meta.Title = strings.TrimSpace(doc.Find(site.Title).First().Text())

	doc.Find(".ttl-grp").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		bind := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, bind, title)
	})

	return meta
}

func get(site Site) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", site.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", site.UserAgent)
	return client.Do(req)
}

func printBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("body: %s", string(body))
}

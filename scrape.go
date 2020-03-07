package main

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
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

	// load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// extract meta data from web page
	meta.Title = site.Title.Text(doc)
	meta.Actor = site.Actor.Text(doc)
	meta.Poster = site.Poster.Image(doc)
	meta.Producer = site.Producer.Text(doc)
	meta.Sample = site.Sample.Link(doc)
	meta.Series = site.Series.Text(doc)
	meta.Release = site.Release.Text(doc)
	meta.Duration = site.Duration.Text(doc)
	meta.Id = site.Id.Text(doc)
	meta.Label = site.Label.Text(doc)
	meta.Genre = site.Genre.Texts(doc)
	meta.Images = site.Images.Images(doc)

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

func printHtmlBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("body: %s", string(body))
}

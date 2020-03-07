package main

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	. "tidy/model"
	. "tidy/selector"
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

	body := resp.Body

	// convert none utf-8 web page to utf-8
	if site.Charset != "" {
		body, err = decodeHTMLBody(resp.Body, site.Charset)
		if err != nil {
			log.Fatal(err)
		}
	}

	// load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
	}

	// extract meta data from web page
	meta.Title = site.Title.Value(doc)
	meta.Actor = site.Actor.Value(doc)
	meta.Poster = site.Poster.Value(doc)
	meta.Producer = site.Producer.Value(doc)
	meta.Sample = site.Sample.Value(doc)
	meta.Series = site.Series.Value(doc)
	meta.Release = site.Release.Value(doc)
	meta.Duration = site.Duration.Value(doc)
	meta.Id = site.Id.Value(doc)
	meta.Label = site.Label.Value(doc)
	meta.Genre = site.Genre.Values(doc)
	meta.Images = site.Images.Values(doc)
	meta.Url = site.Url

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

func decodeHTMLBody(body io.Reader, encoding string) (io.ReadCloser, error) {

	body, err := charset.NewReaderLabel(encoding, body)

	if err != nil {
		log.Fatal(err)
	}

	return ioutil.NopCloser(body), nil
}

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, ""); ok {
			return name
		}
	}
	return "utf-8"
}

func printHtmlBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("body: %s", string(body))
}

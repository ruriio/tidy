package selector

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	. "tidy/model"
)

const UserAgent string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/75.0.3770.90 Safari/537.36"

const MobileUserAgent string = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) " +
	"AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

type Site struct {
	Url       string
	UserAgent string
	Charset   string
	Cookies   []http.Cookie
	CssSelector
	Next *Site
}

func (site Site) Meta() Meta {
	var meta = Meta{}
	// load the HTML document
	doc, err := goquery.NewDocumentFromReader(site.Body())

	if err != nil {
		log.Fatal(err)
	}

	var next = Meta{}
	if site.Next != nil {
		next = site.Next.Meta()
	}

	// extract meta data from web page
	meta.Title = combine(site.Title.Value(doc), next.Title)
	meta.Actor = combine(site.Actor.Value(doc), next.Actor)
	meta.Poster = combine(site.Poster.Value(doc), next.Poster)
	meta.Producer = combine(site.Producer.Value(doc), next.Producer)
	meta.Sample = combine(site.Sample.Value(doc), next.Sample)
	meta.Series = combine(site.Series.Value(doc), next.Series)
	meta.Release = combine(site.Release.Value(doc), next.Release)
	meta.Duration = combine(site.Duration.Value(doc), next.Duration)
	meta.Id = combine(site.Id.Value(doc), next.Id)
	meta.Label = combine(site.Label.Value(doc), next.Label)
	meta.Genre = combineArray(site.Genre.Values(doc), next.Genre)
	meta.Images = combineArray(site.Images.Values(doc), next.Images)
	meta.Url = site.Url
	return meta
}

func combine(first string, second string) string {
	if len(first) > 0 {
		return first
	} else {
		return second
	}
}

func combineArray(first []string, second []string) []string {
	if len(first) > 0 {
		return first
	} else {
		return second
	}
}

func (site Site) Body() io.ReadCloser {
	resp, err := site.get()

	if err != nil {
		log.Fatal(err)
	}
	//defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("stats code error: %d %s", resp.StatusCode, resp.Status)
	}

	body := resp.Body

	//printHtmlBody(resp)

	// convert none utf-8 web page to utf-8
	if site.Charset != "" {
		body, err = decodeHTMLBody(resp.Body, site.Charset)
		if err != nil {
			log.Fatal(err)
		}
	}
	return body
}

func (site Site) get() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", site.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", site.UserAgent)

	for _, cookie := range site.Cookies {
		req.AddCookie(&cookie)
	}

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

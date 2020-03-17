package selector

import (
	"bufio"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	. "github.com/ruriio/tidy/model"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const UserAgent string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/75.0.3770.90 Safari/537.36"

const MobileUserAgent string = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) " +
	"AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

type Site struct {
	Key     string
	Url       string
	WebUrl    string
	UserAgent string
	Charset   string
	Cookies   []http.Cookie
	Json      bool
	JsonData  interface{}
	Next      *Site
	Decor     Decor
	meta      Meta
	Path      string

	Selector
}

func (site Site) Decorate(meta *Meta) *Meta {
	return &site.meta
}

func (site Site) Meta() Meta {
	site.meta = Meta{}
	if site.Json {
		site.meta = site.parseJson()
	} else {
		site.meta = site.parseHtml()
	}

	if len(site.WebUrl) > 0 {
		site.meta.Url = site.WebUrl
	} else {
		site.meta.Url = site.Url
	}

	if len(site.Path) > 0 {
		site.meta.Extras = make(map[string]string)
		site.meta.Extras["path"] = site.path(site.meta)
	}

	if site.Decor != nil {
		return *site.Decor.Decorate(&site.meta)
	} else {
		return site.meta
	}
}

func (site Site) parseJson() Meta {
	var meta = Meta{}
	body, err := ioutil.ReadAll(site.Body())
	err = json.Unmarshal(body, &site.JsonData)
	if err != nil {
		log.Println(err)
	}

	data := make(map[string]interface{})
	m, ok := site.JsonData.(map[string]interface{})
	if ok {
		for k, v := range m {
			//fmt.Println(k, "=>", v)
			data[k] = v
		}
	}

	next := Meta{}
	if site.Next != nil {
		next = site.Next.Meta()
	}

	// extract meta data from json data
	meta.Title = oneOf(site.Title.Query(data), next.Title)
	meta.Actor = oneOf(site.Actor.Query(data), next.Actor)
	meta.Poster = oneOf(site.Poster.Query(data), next.Poster)
	meta.Producer = oneOf(site.Producer.Query(data), next.Producer)
	meta.Sample = oneOf(site.Sample.Query(data), next.Sample)
	meta.Series = oneOf(site.Series.Query(data), next.Series)
	meta.Release = oneOf(site.Release.Query(data), next.Release)
	meta.Duration = oneOf(site.Duration.Query(data), next.Duration)
	meta.Id = oneOf(site.Id.Query(data), next.Id)
	meta.Label = oneOf(site.Label.Query(data), next.Label)
	meta.Genre = oneOfArray(site.Genre.Queries(data), next.Genre)
	meta.Images = oneOfArray(site.Images.Queries(data), next.Images)

	return meta
}

func (site Site) parseHtml() Meta {
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
	meta.Title = oneOf(site.Title.Value(doc), next.Title)
	meta.Actor = oneOf(site.Actor.Value(doc), next.Actor)
	meta.Poster = oneOf(site.Poster.Value(doc), next.Poster)
	meta.Producer = oneOf(site.Producer.Value(doc), next.Producer)
	meta.Sample = oneOf(site.Sample.Value(doc), next.Sample)
	meta.Series = oneOf(site.Series.Value(doc), next.Series)
	meta.Release = oneOf(site.Release.Value(doc), next.Release)
	meta.Duration = oneOf(site.Duration.Value(doc), next.Duration)
	meta.Id = oneOf(site.Id.Value(doc), next.Id)
	meta.Label = oneOf(site.Label.Value(doc), next.Label)
	meta.Genre = oneOfArray(site.Genre.Values(doc), next.Genre)
	meta.Images = oneOfArray(site.Images.Values(doc), next.Images)

	// extract extras to meta
	if site.Extras != nil {
		meta.Extras = make(map[string]string)
		for key, value := range site.Extras {
			meta.Extras[key] = value.Value(doc)
		}

		for key, value := range next.Extras {
			meta.Extras[key] = value
		}
	} else {
		meta.Extras = next.Extras
	}

	return meta
}

func (site Site) Body() io.ReadCloser {
	resp, err := site.get()

	if err != nil {
		log.Println(err)
	}
	//defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("stats code error: %d %s\n", resp.StatusCode, resp.Status)
	}

	body := resp.Body

	//printHtmlBody(resp)

	// convert none utf-8 web page to utf-8
	if site.Charset != "" {
		body, err = decodeHTMLBody(resp.Body, site.Charset)
		if err != nil {
			log.Println(err)
		}
	}
	return body
}

func (site Site) get() (*http.Response, error) {
	log.Printf("url: %s", site.Url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", site.Url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("User-Agent", site.UserAgent)

	for _, cookie := range site.Cookies {
		req.AddCookie(&cookie)
	}

	return client.Do(req)
}

func (site Site) path(meta Meta) string {
	var replacer = strings.NewReplacer("$Title", meta.Title, "$Id", site.Key, "$Actor",
		meta.Actor, "$Series", meta.Series, "$Producer", meta.Producer)
	return replacer.Replace(site.Path)
}

func oneOf(first string, second string) string {
	if len(first) > 0 {
		return first
	} else {
		return second
	}
}

func oneOfArray(first []string, second []string) []string {
	if len(first) > 0 {
		return first
	} else {
		return second
	}
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

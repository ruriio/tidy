package selector

import (
	"bufio"
	"encoding/json"
	"errors"
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
	Key       string
	Url       string
	WebUrl    string
	UserAgent string
	Charset   string
	Cookies   []http.Cookie
	Path      string

	Selector

	Json     bool
	JsonData interface{}

	Decor Decor

	meta   Meta
	Next   *Site
	Search *Site // if directly parse failed, get meta url from search result.
}

func (site Site) Decorate(meta *Meta) *Meta {
	return &site.meta
}

func (site *Site) Meta() Meta {
	site.meta = Meta{}
	if site.Json {
		site.meta = site.parseJson()
	} else {
		site.meta = site.parseHtml()
	}

	if site.meta.Extras == nil {
		site.meta.Extras = make(map[string]interface{})
	}

	if len(site.Path) > 0 {
		site.meta.Extras["path"] = site.path(site.meta)
	}

	if len(site.WebUrl) > 0 {
		site.meta.Url = site.WebUrl
	} else {
		site.meta.Url = site.Url
	}

	if site.Next != nil {
		if len(site.Next.Url) > 0 {
			site.meta.Extras["nextUrl"] = site.Next.Url
		}
	}

	if site.Decor != nil {
		return *site.Decor.Decorate(&site.meta)
	} else {
		return site.meta
	}
}

func (site *Site) parseJson() Meta {
	var meta = Meta{}
	html, err := site.Body()

	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(html)
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

func (site *Site) parseHtml() Meta {
	var meta = Meta{}
	var doc *goquery.Document
	body, err := site.Body()

	if err != nil {
		log.Println(err)
	} else {
		// load the HTML document
		doc, err = goquery.NewDocumentFromReader(body)

		if err != nil {
			log.Println(err)
		}
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

	if meta.Id == "" && site.Next != nil {
		meta.Id = oneOf(site.Key, site.Next.Key)
	}

	// extract extras to meta
	if site.Extras != nil {
		meta.Extras = make(map[string]interface{})
		for key, value := range site.Extras {
			if value.Plural {
				meta.Extras[key] = value.Values(doc)
			} else {
				meta.Extras[key] = value.Value(doc)
			}
		}

		for key, value := range next.Extras {
			meta.Extras[key] = value
		}
	} else {
		meta.Extras = next.Extras
	}

	return meta
}

func (site *Site) Body() (io.ReadCloser, error) {

	if site.Url == "" {
		return site.searchAndGet()
	}

	resp, err := site.get()

	if err != nil {
		log.Println(err)
	}
	//defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("stats code error: %d %s, using search\n", resp.StatusCode, resp.Status)
		return site.searchAndGet()
	} else {
		body := resp.Body

		//printHtmlBody(resp)

		// convert none utf-8 web page to utf-8
		if site.Charset != "" {
			body, err = decodeHTMLBody(resp.Body, site.Charset)
			if err != nil {
				log.Println(err)
			}
		}
		return body, nil
	}
}

func (site *Site) searchAndGet() (io.ReadCloser, error) {
	// get meta url from search result
	url := site.search()
	if len(url) > 0 {
		site.Url = url
		site.WebUrl = url
		return site.Body()
	} else {
		return nil, errors.New("No metadata found for " + site.Key)
	}
}

func (site *Site) get() (*http.Response, error) {
	log.Printf("get: %s", site.Url)
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

func (site *Site) search() string {
	if site.Search == nil {
		return ""
	}

	body, err := site.Search.Body()
	if err != nil {
		log.Println(err)
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Println(err)
		return ""
	}

	hrefs := site.Search.Extras["search"].Values(doc)

	for _, href := range hrefs {
		matcher := site.Search.Extras["match"]
		if matcher != nil {
			if len(matcher.matcherText(href)) > 0 {
				log.Println("match search result: " + href)
				return href
			}
		} else if strings.Contains(href, site.Search.Key) {
			log.Println("find search result: " + href)
			return href
		}
	}

	return ""
}

func (site *Site) path(meta Meta) string {
	key := site.Key

	if len(key) == 0 {
		key = meta.Id
	}

	var replacer = strings.NewReplacer("$Title", meta.Title, "$Id", key,
		"$Actor", oneOf(meta.Actor, meta.Producer, meta.Series),
		"$Series", oneOf(meta.Series, meta.Producer, meta.Actor),
		"$Producer", oneOf(meta.Producer, meta.Series, meta.Actor))
	path := replacer.Replace(site.Path)

	// fix for filename too long error
	if len(path) > 204 {
		path = string([]rune(path)[0:80])
	}

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return path
}

func oneOf(str ...string) string {
	for _, s := range str {
		if len(s) > 0 {
			return s
		}
	}
	return str[0]
}

func oneOfArray(arr ...[]string) []string {
	for _, a := range arr {
		if len(a) > 0 {
			return a
		}
	}
	return arr[0]
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

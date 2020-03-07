package sites

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const UserAgent string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/75.0.3770.90 Safari/537.36"

const MobileUserAgent string = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) " +
	"AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

type Site struct {
	Url       string
	UserAgent string
	Selector
}

type Selector struct {
	Id       Item
	Title    Item
	Actor    Item
	Poster   Item
	Series   Item
	Producer Item
	Release  Item
	Duration Item
	Sample   Item
	Images   Item
	Label    Item
	Genre    Item
}

type Item struct {
	selector string
	replacer *strings.Replacer
}

func selector(selector string) Item {
	return Item{selector: selector, replacer: strings.NewReplacer("", "")}
}

func replacer(selector string, replacer *strings.Replacer) Item {
	return Item{selector: selector, replacer: replacer}
}

func (selector Item) Text(doc *goquery.Document) string {
	text := strings.TrimSpace(doc.Find(selector.selector).First().Text())
	return selector.replacer.Replace(text)
}

func (selector Item) Texts(doc *goquery.Document) []string {
	var texts []string
	doc.Find(selector.selector).Each(func(i int, selection *goquery.Selection) {
		text := strings.TrimSpace(selection.Text())
		text = selector.replacer.Replace(text)
		texts = append(texts, text)
	})

	return texts
}

func (selector Item) Image(doc *goquery.Document) string {
	return selector.Attr(doc, "src")
}

func (selector Item) Images(doc *goquery.Document) []string {
	return selector.Attrs(doc, "src")
}

func (selector Item) Link(doc *goquery.Document) string {
	return selector.Attr(doc, "href")
}

func (selector Item) Attr(doc *goquery.Document, attr string) string {
	src, exist := doc.Find(selector.selector).First().Attr(attr)
	if exist {
		return selector.replacer.Replace(strings.TrimSpace(src))
	}
	return ""
}

func (selector Item) Attrs(doc *goquery.Document, attr string) []string {
	var attrs []string
	doc.Find(selector.selector).Each(func(i int, selection *goquery.Selection) {
		src, exist := selection.Attr(attr)
		if exist {
			text := strings.TrimSpace(src)
			text = selector.replacer.Replace(text)
			attrs = append(attrs, text)
		}
	})
	return attrs
}

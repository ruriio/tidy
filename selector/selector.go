package selector

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

type CssSelector struct {
	Id       *Item
	Title    *Item
	Actor    *Item
	Poster   *Item
	Series   *Item
	Producer *Item
	Release  *Item
	Duration *Item
	Sample   *Item
	Images   *Item
	Label    *Item
	Genre    *Item
	Extras   map[string]*Item
}

type Item struct {
	selector  string
	attribute string
	replacer  *strings.Replacer
	preset    string
	matcher   string
}

func Selector(selector string) *Item {
	return &Item{selector: selector, attribute: "", replacer: strings.NewReplacer("", ""), preset: ""}
}

func Preset(preset string) *Item {
	return &Item{selector: "", attribute: "", replacer: strings.NewReplacer("", ""), preset: preset}
}

func (selector Item) Replace(oldNew ...string) *Item {
	selector.replacer = strings.NewReplacer(oldNew...)
	return &selector
}

func (selector Item) Attribute(attr string) *Item {
	selector.attribute = attr
	return &selector
}

func (selector Item) Match(matcher string) *Item {
	selector.matcher = matcher
	return &selector
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

func (selector *Item) Value(doc *goquery.Document) string {

	if selector == nil {
		return ""
	}

	if len(selector.preset) > 0 {
		return selector.preset
	}

	if len(selector.matcher) > 0 {
		return selector.matcherValue(doc)
	}

	selection := doc.Find(selector.selector).First()
	return selector.textOrAttr(selection)
}

func (selector Item) matcherValue(doc *goquery.Document) string {
	re := regexp.MustCompile(selector.matcher)
	text := doc.Text()

	matches := re.FindAllString(text, -1)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func (selector *Item) Values(doc *goquery.Document) []string {
	var texts []string

	if selector == nil {
		return texts
	}

	doc.Find(selector.selector).Each(func(i int, selection *goquery.Selection) {
		texts = append(texts, selector.textOrAttr(selection))
	})

	return texts
}

func (selector Item) textOrAttr(selection *goquery.Selection) string {
	text := ""
	if len(selector.attribute) > 0 {
		src, exist := selection.Attr(selector.attribute)
		if exist {
			text = src
		}
	} else {
		text = selection.Text()
	}
	return selector.replacer.Replace(strings.TrimSpace(text))
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

func (selector Item) Links(doc *goquery.Document) []string {
	return selector.Attrs(doc, "href")
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
func (selectors CssSelector) AddExtra(key string, selector *Item) CssSelector {
	if selectors.Extras == nil {
		selectors.Extras = make(map[string]*Item)
	}
	selectors.Extras[key] = selector
	return selectors
}

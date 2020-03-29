package selector

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type Selector struct {
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
	format    string
	presets   []string
	matcher   string
	query     string
	Plural    bool
}

func Select(selector string) *Item {
	return &Item{selector: selector, attribute: "", replacer: strings.NewReplacer("", ""), preset: ""}
}

func Selects(selector string) *Item {
	return &Item{selector: selector, replacer: strings.NewReplacer("", ""), Plural: true}
}

func Preset(preset string) *Item {
	return &Item{preset: preset}
}

func Presets(presets []string) *Item {
	return &Item{presets: presets}
}

func Match(matcher string) *Item {
	return &Item{matcher: matcher}
}

func Query(query string) *Item {
	return &Item{query: query}
}

func (selector Item) Replace(oldNew ...string) *Item {
	selector.replacer = strings.NewReplacer(oldNew...)
	return &selector
}

func (selector Item) Format(format string) *Item {
	selector.format = format
	return &selector
}

func (selector Item) Attribute(attr string) *Item {
	selector.attribute = attr
	return &selector
}

func (selector Item) Text(doc *goquery.Document) string {
	text := doc.Find(selector.selector).First().Text()
	return strings.TrimSpace(selector.replacer.Replace(text))
}

func (selector Item) Texts(doc *goquery.Document) []string {
	var texts []string
	doc.Find(selector.selector).Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		text = strings.TrimSpace(selector.replacer.Replace(text))
		texts = append(texts, text)
	})

	return texts
}

func (selector *Item) Value(doc *goquery.Document) string {

	if selector == nil || doc == nil {
		return ""
	}

	if len(selector.preset) > 0 {
		return selector.preset
	}

	if len(selector.matcher) > 0 {
		return selector.matcherValue(doc)
	}

	selection := doc.Find(selector.selector).First()

	value := selector.textOrAttr(selection)

	if len(selector.format) > 0 {
		value = fmt.Sprintf(selector.format, value)
	}
	return value
}

func (selector Item) matcherValue(doc *goquery.Document) string {
	text := doc.Text()
	return selector.matcherText(text)
}

func (selector Item) matcherText(text string) string {
	re := regexp.MustCompile(selector.matcher)

	matches := re.FindAllString(text, -1)
	if len(matches) > 0 {
		text = matches[0]
	} else {
		text = ""
	}
	if selector.replacer != nil {
		text = selector.replacer.Replace(strings.TrimSpace(text))
	}
	return strings.TrimSpace(text)
}

func (selector *Item) Values(doc *goquery.Document) []string {
	var texts []string

	if selector == nil || doc == nil {
		return texts
	}

	if selector.presets != nil {
		return selector.presets
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

	value := strings.TrimSpace(selector.replacer.Replace(text))

	if len(selector.format) > 0 {
		value = fmt.Sprintf(selector.format, value)
	}

	return value
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

func (selector *Item) Query(data map[string]interface{}) string {
	if selector == nil {
		return ""
	}
	if len(selector.preset) > 0 {
		return selector.preset
	}
	return query(data, selector.query)
}

func query(data map[string]interface{}, key string) string {
	value := data[key]

	if value != nil {
		return fmt.Sprintf("%v", value)
	}
	return ""
}

func (selector *Item) Queries(data map[string]interface{}) []string {

	if selector == nil {
		return []string{}
	}

	if selector.presets != nil {
		return selector.presets
	}

	return queries(data, selector.query)
}

func queries(data map[string]interface{}, key string) []string {
	var res []string
	x := data[key]
	if x != nil {
		// if json object is not slice then ignore
		if reflect.ValueOf(x).Kind() == reflect.Slice {
			array := x.([]interface{})
			for _, v := range array {
				var value string

				if reflect.ValueOf(v).Kind() == reflect.Map {
					out, err := json.Marshal(v)
					if err != nil {
						log.Fatal(err)
					}
					value = string(out)
				} else {
					value = fmt.Sprintf("%v", v)
				}
				res = append(res, value)
			}
		}
	}
	return res
}

func (selectors Selector) Extra(key string, selector *Item) Selector {
	if selectors.Extras == nil {
		selectors.Extras = make(map[string]*Item)
	}
	selectors.Extras[key] = selector
	return selectors
}

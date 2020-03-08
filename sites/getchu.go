package sites

import (
	"fmt"
	"net/http"
	. "tidy/selector"
)

func Getchu(id string) Site {
	return Site{
		Url:       fmt.Sprintf("http://dl.getchu.com/i/item%s", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adult_check_flag", Value: "1"}},
		Charset:   "euc-jp",

		CssSelector: CssSelector{
			Title:    Selector("meta[property=\"og:title\"]").Attribute("content"),
			Actor:    Selector("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Poster:   Selector("meta[property=\"og:image\"]").Attribute("content"),
			Producer: Selector("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Sample:   Selector("a[href^=\"http://file.dl.getchu.com/download_sample_file.php\"]").Attribute("href"),
			Series:   Selector("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Release:  Selector("~").Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Selector("~").Match(`動画.*分`),
			Id:       Selector("input[name=id]").Attribute("value"),
			Label:    Selector("null"),
			Genre:    Selector(".item-key > a"),
			Images:   Selector("a[href^=\"/data/item_img\"]").Attribute("href").Replace("/data", "http://dl.getchu.com/data"),
		},
	}
}

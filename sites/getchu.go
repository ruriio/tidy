package sites

import (
	"fmt"
	"net/http"
	. "github.com/ruriio/tidy/selector"
)

func Getchu(id string) Site {
	return Site{
		Url:       fmt.Sprintf("http://dl.getchu.com/i/item%s", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adult_check_flag", Value: "1"}},
		Charset:   "euc-jp",

		Selector: Selector{
			Title:    Select("meta[property=\"og:title\"]").Attribute("content"),
			Actor:    Select("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Poster:   Select("meta[property=\"og:image\"]").Attribute("content"),
			Producer: Select("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Sample:   Select("a[href^=\"http://file.dl.getchu.com/download_sample_file.php\"]").Attribute("href"),
			Series:   Select("a[href^=\"http://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Match(`動画.*分`),
			Id:       Select("input[name=id]").Attribute("value"),
			Label:    Select("null"),
			Genre:    Select(".item-key > a"),
			Images:   Select("a[href^=\"/data/item_img\"]").Attribute("href").Replace("/data", "http://dl.getchu.com/data"),
		},
	}
}

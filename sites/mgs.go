package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"net/http"
)

func Mgs(id string) Site {

	mobile := Site{
		Url:       fmt.Sprintf("https://sp.mgstage.com/product/product_detail/SP-%s/", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adc", Value: "1"}},

		Selector: Selector{
			Title:  Select(".sample-image-wrap.h1 > img").Attribute("alt"),
			Actor:  Select("a.actor"),
			Poster: Select(".sample-image-wrap.h1").Attribute("href"),
			Sample: Select("#sample-movie").Attribute("src"),
			Series: Select("a.series"),
			Label:  Select("null"),
			Images: Select("a[class^=\"sample-image-wrap sample\"]").Attribute("href"),
		},
	}

	return Site{
		Url:       fmt.Sprintf("https://www.mgstage.com/product/product_detail/%s/", id),
		UserAgent: UserAgent,
		Cookies:   mobile.Cookies,
		Selector: Selector{
			Producer: Select("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(2) > td > a"),
			Release:  Select("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(5) > td"),
			Duration: Select("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(3) > td"),
			Id:       Select("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(4) > td"),
			Genre:    Select("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(9) > td > a"),
		},
		Next: &mobile,
	}
}

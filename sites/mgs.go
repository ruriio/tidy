package sites

import (
	"fmt"
	"net/http"
	. "tidy/selector"
)

func Mgs(id string) Site {

	mobile := Site{
		Url:       fmt.Sprintf("https://sp.mgstage.com/product/product_detail/SP-%s/", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adc", Value: "1"}},

		CssSelector: CssSelector{
			Title:  Selector(".sample-image-wrap.h1 > img").Attribute("alt"),
			Actor:  Selector("a.actor"),
			Poster: Selector(".sample-image-wrap.h1").Attribute("href"),
			Sample: Selector("#sample-movie").Attribute("src"),
			Series: Selector("a.series"),
			Label:  Selector("null"),
			Images: Selector("a[class^=\"sample-image-wrap sample\"]").Attribute("href"),
		},
	}

	return Site{
		Url:       fmt.Sprintf("https://www.mgstage.com/product/product_detail/%s/", id),
		UserAgent: UserAgent,
		Cookies:   mobile.Cookies,
		CssSelector: CssSelector{
			Producer: Selector("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(2) > td > a"),
			Release:  Selector("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(5) > td"),
			Duration: Selector("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(3) > td"),
			Id:       Selector("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(4) > td"),
			Genre:    Selector("div.detail_left > div > table:nth-child(3) > tbody > tr:nth-child(9) > td > a"),
		},
		Next: &mobile,
	}
}

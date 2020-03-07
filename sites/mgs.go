package sites

import (
	"fmt"
	"net/http"
	. "tidy/selector"
)

func Mgs(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://sp.mgstage.com/product/product_detail/SP-%s/", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adc", Value: "1"}},

		CssSelector: CssSelector{
			Title:    Selector(".sample-image-wrap.h1 > img").Attribute("alt"),
			Actor:    Selector("a.actor"),
			Poster:   Selector(".sample-image-wrap.h1").Attribute("href"),
			Producer: Selector("#detail > div > dl > dd:nth-child(6) > a"),
			Sample:   Selector("#sample-movie").Attribute("src"),
			Series:   Selector("a.series"),
			Release:  Selector("#detail > div > dl > dd:nth-child(12)"),
			Duration: Selector("#detail > div > dl > dd:nth-child(14)"),
			Id:       Selector("#detail > div > dl > dd:nth-child(16)").Replace("SP-", ""),
			Label:    Selector("null"),
			Genre:    Selector("#detail > div > dl > dd > a"),
			Images:   Selector("a[class^=\"sample-image-wrap sample\"]").Attribute("href"),
		},
	}
}

package sites

import (
	"fmt"
	. "tidy/selector"
)

func Tokyo(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://my.tokyo-hot.com/product/%s/", id),
		UserAgent: MobileUserAgent,

		CssSelector: CssSelector{
			Title:    Selector(".pagetitle"),
			Actor:    Selector("div.infowrapper > dl > dd:nth-child(2) > a"),
			Poster:   Selector("video").Attribute("poster").Replace("410x231", "820x462"),
			Producer: Preset("Tokyo"),
			Sample:   Selector("source").Attribute("src"),
			Series:   Selector("div.infowrapper > dl > dd > a[href^=\"/product/?type=genre\"]"),
			Release:  Matcher(`\d{4}/\d{2}/\d{2}`),
			Duration: Matcher(`\d{2}:\d{2}:\d{2}`),
			Id:       Selector("input[name=\"product_uid\"]").Attribute("value"),
			Label:    Selector("div.infowrapper > dl > dd > a[href^=\"/product/?type=vendor\"], div.infowrapper > dl > dd > a[href^=\"/product/?vendor\"]"),
			Genre:    Selector("div.infowrapper > dl > dd > a[href^=\"/product/?type=tag\"]"),
			Images:   Selector("a[rel=\"cap\"]").Attribute("href").Replace(" ", "%20"),
		},
	}
}

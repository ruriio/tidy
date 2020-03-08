package sites

import (
	"fmt"
	. "tidy/selector"
)

func Tokyo(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://my.tokyo-hot.com/product/%s/", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    Select(".pagetitle"),
			Actor:    Select("div.infowrapper > dl > dd:nth-child(2) > a"),
			Poster:   Select("video").Attribute("poster").Replace("410x231", "820x462"),
			Producer: Preset("Tokyo"),
			Sample:   Select("source").Attribute("src"),
			Series:   Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=genre\"]"),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Match(`\d{2}:\d{2}:\d{2}`),
			Id:       Select("input[name=\"product_uid\"]").Attribute("value"),
			Label:    Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=vendor\"], div.infowrapper > dl > dd > a[href^=\"/product/?vendor\"]"),
			Genre:    Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=tag\"]"),
			Images:   Select("a[rel=\"cap\"]").Attribute("href").Replace(" ", "%20"),
		},
	}
}

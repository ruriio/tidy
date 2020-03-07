package sites

import (
	"fmt"
	. "tidy/selector"
)

func Dmm(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", id),
		UserAgent: MobileUserAgent,

		CssSelector: CssSelector{
			Title:    Selector(".ttl-grp"),
			Actor:    Selector("ul.parts-maindata > li > a > span"),
			Poster:   Selector(".package").Replace("ps.jpg", "pl.jpg").Attribute("src"),
			Producer: Selector(".parts-subdata"),
			Sample:   Selector(".play-btn"),
			Series:   Selector("#work-mono-info > dl:nth-child(4) > dd"),
			Release:  Selector("#work-mono-info > dl:nth-child(8) > dd"),
			Duration: Selector("#work-mono-info > dl:nth-child(9) > dd"),
			Id:       Selector("#work-mono-info > dl:nth-child(10) > dd"),
			Label:    Selector("#work-mono-info > dl:nth-child(6) > dd > ul > li > a"),
			Genre:    Selector("#work-mono-info > dl.box-genreinfo > dd > ul > li > a"),
			Images:   Selector("#sample-list > ul > li > a > span > img").Replace("-", "jp-").Attribute("src"),
		},
	}
}

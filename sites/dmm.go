package sites

import (
	"fmt"
	. "strings"
)

func Dmm(id string) Site {
	return Site{
		Url:       url(id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    selector(".ttl-grp"),
			Actor:    selector("ul.parts-maindata > li > a > span"),
			Poster:   replacer(".package", NewReplacer("ps.jpg", "pl.jpg")),
			Producer: selector(".parts-subdata"),
			Sample:   selector(".play-btn"),
			Series:   selector("#work-mono-info > dl:nth-child(4) > dd"),
			Release:  selector("#work-mono-info > dl:nth-child(8) > dd"),
			Duration: selector("#work-mono-info > dl:nth-child(9) > dd"),
			Id:       selector("#work-mono-info > dl:nth-child(10) > dd"),
			Label:    selector("#work-mono-info > dl:nth-child(6) > dd > ul > li > a"),
			Genre:    selector("#work-mono-info > dl.box-genreinfo > dd > ul > li > a"),
			Images:   replacer("#sample-list > ul > li > a > span > img", NewReplacer("-", "jp-")),
		},
	}
}

func url(id string) string {
	return fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", id)
}

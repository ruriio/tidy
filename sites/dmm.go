package sites

import (
	"fmt"
)

func Dmm(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    selector(".ttl-grp"),
			Actor:    selector("ul.parts-maindata > li > a > span"),
			Poster:   selector(".package").replace("ps.jpg", "pl.jpg").attr("src"),
			Producer: selector(".parts-subdata"),
			Sample:   selector(".play-btn"),
			Series:   selector("#work-mono-info > dl:nth-child(4) > dd"),
			Release:  selector("#work-mono-info > dl:nth-child(8) > dd"),
			Duration: selector("#work-mono-info > dl:nth-child(9) > dd"),
			Id:       selector("#work-mono-info > dl:nth-child(10) > dd"),
			Label:    selector("#work-mono-info > dl:nth-child(6) > dd > ul > li > a"),
			Genre:    selector("#work-mono-info > dl.box-genreinfo > dd > ul > li > a"),
			Images:   selector("#sample-list > ul > li > a > span > img").replace("-", "jp-").attr("src"),
		},
	}
}

package sites

import (
	"fmt"
	. "tidy/selector"
)

func CaribPr(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://www.caribbeancompr.com/moviepages/%s/index.html", id),
		UserAgent: MobileUserAgent,
		Charset:   "euc-jp",

		Selector: Selector{
			Title:    Select("h1"),
			Actor:    Select("a.spec-item[href^=\"/search/\"]"),
			Poster:   Preset(fmt.Sprintf("https://www.caribbeancompr.com/moviepages/%s/images/l_l.jpg", id)),
			Producer: Preset("Caribbean"),
			Sample:   Preset(fmt.Sprintf("https://smovie.caribbeancompr.com/sample/movies/%s/480p.mp4", id)),
			Series:   Select("a[href^=\"/serieslist/\"]"),
			Release:  Select("div.movie-info > div > ul > li:nth-child(2) > span.spec-content"),
			Duration: Select("div.movie-info > div > ul > li:nth-child(3) > span.spec-content"),
			Id:       Preset(id),
			Label:    Select("a[href^=\"/serieslist/\"]"),
			Genre:    Select("a.spec-item[href^=\"/listpages/\"]"),
			Images:   Select("a[data-is_sample='1']").Attribute("href"),
		},
	}
}

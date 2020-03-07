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

		CssSelector: CssSelector{
			Title:    Selector("h1"),
			Actor:    Selector("a.spec-item[href^=\"/search/\"]"),
			Poster:   Preset(fmt.Sprintf("https://www.caribbeancompr.com/moviepages/%s/images/l_l.jpg", id)),
			Producer: Preset("Caribbean"),
			Sample:   Preset(fmt.Sprintf("https://smovie.caribbeancompr.com/sample/movies/%s/480p.mp4", id)),
			Series:   Selector("a[href^=\"/serieslist/\"]"),
			Release:  Selector("div.movie-info > div > ul > li:nth-child(2) > span.spec-content"),
			Duration: Selector("div.movie-info > div > ul > li:nth-child(3) > span.spec-content"),
			Id:       Preset(id),
			Label:    Selector("a[href^=\"/serieslist/\"]"),
			Genre:    Selector("a.spec-item[href^=\"/listpages/\"]"),
			Images:   Selector("a[data-is_sample='1']").Attribute("href"),
		},
	}
}

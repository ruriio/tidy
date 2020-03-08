package sites

import (
	"fmt"
	. "tidy/selector"
)

func Carib(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://www.caribbeancom.com/moviepages/%s/index.html", id),
		UserAgent: MobileUserAgent,
		Charset:   "euc-jp",

		Selector: Selector{
			Title:    Select("h1[itemprop=name]"),
			Actor:    Select("a[itemprop=actor]"),
			Poster:   Preset(fmt.Sprintf("https://www.caribbeancom.com/moviepages/%s/images/l_l.jpg", id)),
			Producer: Preset("Caribbean"),
			Sample:   Preset(fmt.Sprintf("https://smovie.caribbeancom.com/sample/movies/%s/480p.mp4", id)),
			Series:   Select("a[onclick^=gaDetailEvent\\(\\'Series\\ Name\\']"),
			Release:  Select("span[itemprop=datePublished]"),
			Duration: Select("span[itemprop=duration]"),
			Id:       Preset(id),
			Label:    Select("null"),
			Genre:    Select("a[itemprop=genre]"),
			Images:   Select("a[data-is_sample='1']").Attribute("href").Replace("/movie", "https://www.caribbeancom.com/movie"),
		},
	}
}

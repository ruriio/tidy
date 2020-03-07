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

		CssSelector: CssSelector{
			Title:    Selector("h1[itemprop=name]"),
			Actor:    Selector("a[itemprop=actor]"),
			Poster:   Preset(fmt.Sprintf("https://www.caribbeancom.com/moviepages/%s/images/l_l.jpg", id)),
			Producer: Preset("Caribbean"),
			Sample:   Preset(fmt.Sprintf("https://smovie.caribbeancom.com/sample/movies/%s/480p.mp4", id)),
			Series:   Selector("a[onclick^=gaDetailEvent\\(\\'Series\\ Name\\']"),
			Release:  Selector("span[itemprop=datePublished]"),
			Duration: Selector("span[itemprop=duration]"),
			Id:       Preset(id),
			Label:    Selector("null"),
			Genre:    Selector("a[itemprop=genre]"),
			Images:   Selector("a[data-is_sample='1']").Attribute("href").Replace("/movie", "https://www.caribbeancom.com/movie"),
		},
	}
}

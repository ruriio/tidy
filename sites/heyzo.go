package sites

import (
	"fmt"
	. "tidy/selector"
)

func Heyzo(id string) Site {
	return Site{
		Url:       fmt.Sprintf("http://m.heyzo.com/moviepages/%s/index.html", id),
		UserAgent: MobileUserAgent,

		CssSelector: CssSelector{
			Title:    Selector("h1"),
			Actor:    Selector("strong.name"),
			Poster:   Selector("#gallery > div > a > img").Attribute("src").Replace("gallery/thumbnail_001.jpg", "images/player_thumbnail.jpg"),
			Producer: Preset("HEYZO"),
			Sample:   Selector("#gallery > div > a > img").Attribute("src").Replace("gallery/thumbnail_001.jpg", "sample.mp4"),
			Series:   Selector("#series").Replace("シリーズ：", ""),
			Release:  Selector("span.release").Replace("配信日：", ""),
			Duration: Selector("span[itemprop=duration]"),
			Id:       Selector("input[name=movie_id]").Attribute("value"),
			Label:    Selector("null"),
			Genre:    Selector("#keyword > ul > ul > li > a"),
			Images:   Selector("#gallery > div > a > img").Attribute("src").Replace("thumbnail_", ""),
		},
	}
}

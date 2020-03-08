package sites

import (
	"fmt"
	. "tidy/selector"
)

func Heyzo(id string) Site {
	const providerId = "provider_id"

	return Site{
		Url:       fmt.Sprintf("http://m.heyzo.com/moviepages/%s/index.html", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    Select("h1"),
			Actor:    Select("strong.name"),
			Poster:   Select("#gallery > div > a > img").Attribute("src").Replace("gallery/thumbnail_001.jpg", "images/player_thumbnail.jpg"),
			Producer: Preset("HEYZO"),
			Sample:   Select("#gallery > div > a > img").Attribute("src").Replace("gallery/thumbnail_001.jpg", "sample.mp4"),
			Series:   Select("#series").Replace("シリーズ：", ""),
			Release:  Select("span.release").Replace("配信日：", ""),
			Duration: Select("span[itemprop=duration]"),
			Id:       Select("input[name=movie_id]").Attribute("value"),
			Label:    Select("null"),
			Genre:    Select("#keyword > ul > ul > li > a"),
			Images:   Select("#gallery > div > a > img").Attribute("src").Replace("thumbnail_", ""),
		}.AddExtra(providerId, Select("input[name=provider_id]").Attribute("value")),
	}
}

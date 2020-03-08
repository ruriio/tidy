package sites

import (
	"fmt"
	. "tidy/selector"
)

func Fc2Club(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://fc2club.com/html/FC2-%s.html", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    Select("div.col-sm-8 > h3").Replace(fmt.Sprintf("FC2-%s ", id), ""),
			Actor:    Select("div.col-sm-8 > h5:nth-child(7) > a"),
			Poster:   Select("div.col-sm-8 > a").Attribute("href").Replace("/upload", "https://fc2club.com/upload"),
			Producer: Select("div.col-sm-8 > h5:nth-child(5) > a:nth-child(2)"),
			Sample:   Select("null"),
			Series:   Select("div.col-sm-8 > h5:nth-child(5) > a:nth-child(2)"),
			Release:  Select(".items_article_Releasedate").Replace("販売日 : ", ""),
			Duration: Select(".items_article_MainitemThumb > p"),
			Id:       Select(".items_article_TagArea").Attribute("data-id"),
			Label:    Select("null"),
			Genre:    Select("null"),
			Images:   Select("ul.slides > li > img").Attribute("src").Replace("/upload", "https://fc2club.com/upload"),
		},
	}
}

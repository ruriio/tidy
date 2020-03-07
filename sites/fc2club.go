package sites

import (
	"fmt"
	. "tidy/selector"
)

func Fc2Club(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://fc2club.com/html/FC2-%s.html", id),
		UserAgent: MobileUserAgent,

		CssSelector: CssSelector{
			Title:    Selector("div.col-sm-8 > h3").Replace(fmt.Sprintf("FC2-%s ", id), ""),
			Actor:    Selector("div.col-sm-8 > h5:nth-child(7) > a"),
			Poster:   Selector("div.col-sm-8 > a").Attribute("href").Replace("/upload", "https://fc2club.com/upload"),
			Producer: Selector("div.col-sm-8 > h5:nth-child(5) > a:nth-child(2)"),
			Sample:   Selector("null"),
			Series:   Selector("div.col-sm-8 > h5:nth-child(5) > a:nth-child(2)"),
			Release:  Selector(".items_article_Releasedate").Replace("販売日 : ", ""),
			Duration: Selector(".items_article_MainitemThumb > p"),
			Id:       Selector(".items_article_TagArea").Attribute("data-id"),
			Label:    Selector("null"),
			Genre:    Selector("null"),
			Images:   Selector("ul.slides > li > img").Attribute("src").Replace("/upload", "https://fc2club.com/upload"),
		},
	}
}

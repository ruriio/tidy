package sites

import (
	"fmt"
	. "tidy/selector"
)

func Fc2(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://adult.contents.fc2.com/article/%s/", id),
		UserAgent: MobileUserAgent,

		CssSelector: CssSelector{
			Title:    Selector(".items_article_MainitemNameTitle"),
			Actor:    Selector(".items_article_seller").Replace("by ", ""),
			Poster:   Selector("meta[property^=\"og:image\"]").Attribute("content"),
			Producer: Selector(".items_article_seller").Replace("by ", ""),
			Sample:   Selector(".main-video").Attribute("src"),
			Series:   Selector(".items_article_seller").Replace("by ", ""),
			Release:  Selector(".items_article_Releasedate").Replace("販売日 : ", ""),
			Duration: Selector(".items_article_MainitemThumb > p"),
			Id:       Selector(".items_article_TagArea").Attribute("data-id"),
			Label:    Selector("null"),
			Genre:    Selector("null"),
			Images:   Selector("li[data-img^=\"https://storage\"]").Attribute("data-img"),
		},
	}
}

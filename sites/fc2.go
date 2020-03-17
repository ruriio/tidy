package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
)

func Fc2(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://adult.contents.fc2.com/article/%s/", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    Select(".items_article_MainitemNameTitle"),
			Actor:    Select(".items_article_seller").Replace("by ", ""),
			Poster:   Select("meta[property^=\"og:image\"]").Attribute("content"),
			Producer: Select(".items_article_seller").Replace("by ", ""),
			Sample:   Select(".main-video").Attribute("src"),
			Series:   Select(".items_article_seller").Replace("by ", ""),
			Release:  Select(".items_article_Releasedate").Replace("販売日 : ", ""),
			Duration: Select(".items_article_MainitemThumb > p"),
			Id:       Select(".items_article_TagArea").Attribute("data-id"),
			Label:    Select("null"),
			Genre:    Select("null"),
			Images:   Select("li[data-img^=\"https://storage\"]").Attribute("data-img"),
		},
	}
}

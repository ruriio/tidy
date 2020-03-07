package sites

import (
	"fmt"
)

func Fc2(id string) Site {
	return Site{
		Url:       fmt.Sprintf("https://adult.contents.fc2.com/article/%s/", id),
		UserAgent: MobileUserAgent,

		Selector: Selector{
			Title:    selector(".items_article_MainitemNameTitle"),
			Actor:    selector(".items_article_seller").replace("by ", ""),
			Poster:   selector("meta[property^=\"og:image\"]").attr("content"),
			Producer: selector(".items_article_seller").replace("by ", ""),
			Sample:   selector(".main-video").attr("src"),
			Series:   selector(".items_article_seller").replace("by ", ""),
			Release:  selector(".items_article_Releasedate").replace("販売日 : ", ""),
			Duration: selector(".items_article_MainitemThumb > p"),
			Id:       selector(".items_article_TagArea").attr("data-id"),
			Label:    selector("null"),
			Genre:    selector("null"),
			Images:   selector("li[data-img^=\"https://storage\"]").attr("data-img"),
		},
	}
}

package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"regexp"
	"strings"
)

func Fc2(id string) Site {
	id = parseFc2Id(id)
	next := Fc2Club(id)

	return Site{
		Url:       fmt.Sprintf("https://adult.contents.fc2.com/article/%s/", id),
		UserAgent: MobileUserAgent,
		Path:      "fc2/$Actor/FC2-PPV-$Id $Title/",
		Next:      &next,

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

func parseFc2Id(id string) string {
	id = strings.ToLower(id)
	if strings.HasPrefix(id, "fc2") {
		re := regexp.MustCompile(`\d{4,}`)
		matches := re.FindAllString(id, -1)
		if len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}

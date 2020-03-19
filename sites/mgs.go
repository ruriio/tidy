package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"net/http"
	"path"
	"strings"
)

func Mgs(id string) Site {

	id = parseMgsKey(id)

	return Site{
		Url:       fmt.Sprintf("https://sp.mgstage.com/product/product_detail/SP-%s/", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adc", Value: "1"}},
		Path:      "mgs/$Series/$Id $Title/",

		Selector: Selector{
			Id:       Preset(id),
			Title:    Select(".sample-image-wrap.h1 > img").Attribute("alt"),
			Actor:    Select("a.actor"),
			Poster:   Select(".sample-image-wrap.h1").Attribute("href"),
			Sample:   Select("#sample-movie").Attribute("src"),
			Series:   Select("a.series"),
			Label:    Select("null"),
			Images:   Select("a[class^=\"sample-image-wrap sample\"]").Attribute("href"),
			Duration: Match(`\d{2,}分`),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Producer: Match(`メーカー\s.*\s`).Replace("メーカー", ""),
			Genre:    Select(".info > dl > dd > a"),
		},
	}
}

func parseMgsKey(key string) string {
	ext := path.Ext(key)
	return strings.ToUpper(strings.TrimSuffix(key, ext))
}

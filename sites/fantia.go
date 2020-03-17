package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"net/http"
)

func Fantia(id string) Site {

	return Site{
		Url:       fmt.Sprintf("https://fantia.jp/products/%s", id),
		UserAgent: UserAgent,
		Cookies:   []http.Cookie{{Name: "_session_id", Value: "5602e9a9f48bba1997b07baca88e525f"}},

		Selector: Selector{
			Title:    Select(".product-title"),
			Actor:    Select("h3.fanclub-name"),
			Poster:   Select("img[src^=\"https://c.fantia.jp/uploads/product/image\"]").Attribute("src"),
			Producer: Select("h3.fanclub-name"),
			Sample:   Select("null"),
			Series:   Select("h3.fanclub-name"),
			Id:       Select("a.btn.btn-default.btn-sm.btn-star").Attribute("data-product_id"),
			Label:    Select("null"),
			Genre:    Select("null"),
			Images:   Select("img[src^=\"https://c.fantia.jp/uploads/product_image\"]").Attribute("src"),
		},
	}
}

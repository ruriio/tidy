package sites

import (
	"fmt"
	"net/http"
	. "tidy/selector"
)

func Fantia(id string) Site {

	return Site{
		Url:       fmt.Sprintf("https://fantia.jp/products/%s", id),
		UserAgent: UserAgent,
		Cookies:   []http.Cookie{{Name: "_session_id", Value: "5602e9a9f48bba1997b07baca88e525f"}},

		CssSelector: CssSelector{
			Title:    Selector(".product-title"),
			Actor:    Selector("h3.fanclub-name"),
			Poster:   Selector("img[src^=\"https://c.fantia.jp/uploads/product/image\"]").Attribute("src"),
			Producer: Selector("h3.fanclub-name"),
			Sample:   Selector("null"),
			Series:   Selector("h3.fanclub-name"),
			Id:       Selector("a.btn.btn-default.btn-sm.btn-star").Attribute("data-product_id"),
			Label:    Selector("null"),
			Genre:    Selector("null"),
			Images:   Selector("img[src^=\"https://c.fantia.jp/uploads/product_image\"]").Attribute("src"),
		},
	}
}

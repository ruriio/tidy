package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"path"
	"regexp"
	"strings"
)

func Ave(id string) Site {
	id = parseAveId(id)

	search := Site{
		Url: fmt.Sprintf("https://www.aventertainments.com/search_Products.aspx?"+
			"languageID=2&dept_id=29&keyword=%s&searchby=keyword", id),
		UserAgent: UserAgent,
		Selector: Selector{}.Extra("search",
			Select(".list-cover > a").Attribute("href")),
	}

	return Site{
		Url:       "",
		UserAgent: UserAgent,
		Search:    &search,
		Path:      "ave/$Actor/$Id $Title/",

		Selector: Selector{
			Title: Select("h2"),
			Actor: Select("a[href^=\"https://www.aventertainments.com/ActressDetail\"]"),
			Poster: Select("img[src^=\"https://imgs.aventertainments.com/new/jacket_images\"]").
				Attribute("src").Replace("jacket_images", "bigcover"),
			Producer: Select("a[href^=\"https://www.aventertainments.com/studio_products\"]"),
			Sample:   Match(`https://.*.m3u8`),
			Series:   Select("a[href^=\"https://www.aventertainments.com/Series\""),
			Release:  Match(`\d{1,}/\d{1,}/\d{4}`),
			Duration: Match(`\d{2,} Min`),
			Id:       Select(".top-title").Replace("商品番号: ", ""),
			Label:    Select("null"),
			Genre:    Select("ol > a[href^=\"https://www.aventertainments.com/subdept_product\"]"),
			Images: Select("img[src^=\"https://imgs.aventertainments.com/new/screen_shot\"], "+
				"img[src^=\"https://imgs.aventertainments.com//vodimages/screenshot/\"]").
				Attribute("src").Replace("small", "large"),
		}.Extra("actors", Selects("a[href^=\"https://www.aventertainments.com/ActressDetail\"]")),
	}
}

func parseAveId(key string) string {
	ext := path.Ext(key)
	name := strings.ToUpper(strings.TrimSuffix(key, ext))
	re := regexp.MustCompile(`[A-Z]{2,}-? ?\d{2,}`)

	matches := re.FindAllString(name, -1)

	if len(matches) > 0 {
		return matches[0]
	}
	return "nil"
}

package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"golang.org/x/text/encoding/japanese"
	"net/http"
	"net/url"
)

func Getchu(id string) Site {
	sStr, _ := japanese.EUCJP.NewEncoder().String("検索")
	keyStr, _ := japanese.EUCJP.NewEncoder().String(id)
	search := Site{
		Key:     id,
		Cookies: []http.Cookie{{Name: "adult_check_flag", Value: "1"}},
		Charset: "euc-jp",
		Url: fmt.Sprintf("https://dl.getchu.com/search/search_list.php?search_keyword=%s"+
			"&dojin=1&search_category_id=&action=search&btnWordSearch=%s",
			url.QueryEscape(keyStr),
			url.QueryEscape(sStr)),
		UserAgent: MobileUserAgent,
		Selector: Selector{}.
			Extra("search", Select("a[href^=\"https://dl.getchu.com/i/item\"]").Attribute("href")).
			Extra("match", Match(fmt.Sprintf("\\d{5,8}"))),
	}
	return Site{
		Url:       fmt.Sprintf("https://dl.getchu.com/i/item%s", id),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "adult_check_flag", Value: "1"}},
		Charset:   "euc-jp",
		Path:      "getchu/$Actor/ITEM-$Id $Title/",
		Search:    &search,

		Selector: Selector{
			Title:    Select("meta[property=\"og:title\"]").Attribute("content").Replace("/", " "),
			Actor:    Select("a[href^=\"https://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Poster:   Select("meta[property=\"og:image\"]").Attribute("content"),
			Producer: Select("a[href^=\"https://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Sample:   Select("a[href*=\".dl.getchu.com/download_sample_file.php\"]").Attribute("href"),
			Series:   Select("a[href^=\"https://dl.getchu.com/search/dojin_circle_detail.php\"]"),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Match(`動画.*分`),
			Id:       Select("input[name=id]").Attribute("value"),
			Label:    Select("null"),
			Genre:    Select(".item-key > a"),
			Images:   Select("a[href^=\"/data/item_img\"]").Attribute("href").Replace("/data", "http://dl.getchu.com/data"),
		},
	}
}

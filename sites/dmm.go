package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"net/http"
	"regexp"
	"strings"
)

func Dmm(id string) Site {
	dmmId := parseDmmId(id)
	next := Jav(id)
	search := Site{
		Key:       dmmId,
		Url:       fmt.Sprintf("https://www.dmm.co.jp/search/=/searchstr=%s/", dmmId),
		UserAgent: MobileUserAgent,
		Selector: Selector{}.
			Extra("search", Select("a[href^=\"https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=\"]").Attribute("href")).
			Extra("match", Match(fmt.Sprintf("cid=[a-z_]{0,4}\\d{0,4}%s", dmmId))),
	}
	return Site{
		Key:       parseDmmKey(id),
		Url:       fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", dmmId),
		UserAgent: MobileUserAgent,
		Cookies:   []http.Cookie{{Name: "age_check_done", Value: "1"}},
		Path:      "dmm/$Actor/$Id $Title/",
		Search:    &search,
		Next:      &next,

		Selector: Selector{
			Title:    Select("hgroup > h1").Replace("DVD", "", "Blu-ray", ""),
			Actor:    Select("ul.parts-maindata > li > a > span"),
			Poster:   Select(".package").Replace("ps.jpg", "pl.jpg").Attribute("src"),
			Producer: Select(".parts-subdata"),
			Sample:   Select(".play-btn").Attribute("href"),
			Series:   Select(".box-taglink > li > a[href^=\"/mono/dvd/-/list/=/article=series/\"]"),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Match(`\d{2,}分`),
			Id:       Select("品番 [a-z_\\d]*").Replace("品番 ", ""),
			Label:    Select(".box-taglink > li > a[href^=\"/mono/dvd/-/list/=/article=label/\"]"),
			Genre:    Select(".box-taglink > li > a[href^=\"/mono/dvd/-/list/=/article=keyword/\"]"),
			Images:   Select("#sample-list > ul > li > a > span > img").Replace("-", "jp-").Attribute("src"),
		}.Extra("actors", Selects(".box-taglink > li > a[href^=\"/mono/dvd/-/list/=/article=actress/\"]")),
	}
}

func parseDmmKey(key string) string {
	name := strings.ToUpper(key)
	re := regexp.MustCompile(`[A-Z]{2,}-? ?\d{2,}`)

	matches := re.FindAllString(name, -1)

	if len(matches) > 0 {
		return matches[0]
	}
	return "nil"
}

func parseDmmId(key string) string {
	return strings.ReplaceAll(strings.ToLower(parseDmmKey(key)), "-", "")
}

package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"path"
	"regexp"
	"strings"
)

func Dmm(id string) Site {
	dmmId := parseDmmId(id)
	search := Site{
		Key:       dmmId,
		Url:       fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/search/=/searchstr=%s/", dmmId),
		UserAgent: MobileUserAgent,
		Selector: Selector{}.
			Extra("search", Select("a[href^=\"https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=\"]").Attribute("href")).
			Extra("match", Match(fmt.Sprintf("cid=\\d{0,4}%s", dmmId))),
	}
	return Site{
		Key:       parseDmmKey(id),
		Url:       fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", dmmId),
		UserAgent: MobileUserAgent,
		Path:      "dmm/$Actor/$Id $Title/",
		Search:    &search,

		Selector: Selector{
			Title:    Select("hgroup > h1").Replace("DVD", "", "Blu-ray", ""),
			Actor:    Select("ul.parts-maindata > li > a > span"),
			Poster:   Select(".package").Replace("ps.jpg", "pl.jpg").Attribute("src"),
			Producer: Select(".parts-subdata"),
			Sample:   Select(".play-btn").Attribute("href"),
			Series:   Select("#work-mono-info > dl:nth-child(4) > dd"),
			Release:  Select("#work-mono-info > dl:nth-child(8) > dd"),
			Duration: Select("#work-mono-info > dl:nth-child(9) > dd"),
			Id:       Select("#work-mono-info > dl:nth-child(10) > dd"),
			Label:    Select("#work-mono-info > dl:nth-child(6) > dd > ul > li > a"),
			Genre:    Select("#work-mono-info > dl.box-genreinfo > dd > ul > li > a"),
			Images:   Select("#sample-list > ul > li > a > span > img").Replace("-", "jp-").Attribute("src"),
		}.Extra("actors", Selects(".box-taglink > li > a[href^=\"/mono/dvd/-/list/=/article=actress/\"]")),
	}
}

func parseDmmKey(key string) string {
	ext := path.Ext(key)
	name := strings.ToUpper(strings.TrimSuffix(key, ext))
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

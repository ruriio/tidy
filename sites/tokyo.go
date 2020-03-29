package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
	"path"
	"regexp"
	"strings"
)

func Tokyo(id string) Site {
	id = parseTokyoKey(id)

	search := Site{
		Url:       fmt.Sprintf("https://my.tokyo-hot.com/product/?q=%s&lang=jp", id),
		UserAgent: MobileUserAgent,
		Selector: Selector{}.Extra("search", Select("a.rm").Attribute("href").
			Format("https://my.tokyo-hot.com/%s?lang=jp")),
	}

	return Site{
		Url:       fmt.Sprintf("https://my.tokyo-hot.com/product/%s/?lang=jp", id),
		UserAgent: MobileUserAgent,
		Path:      "tokyo/$Actor/$Id $Title/",
		Search:    &search,

		Selector: Selector{
			Title:    Select(".pagetitle"),
			Actor:    Select("div.infowrapper > dl > dd:nth-child(2) > a"),
			Poster:   Select("video").Attribute("poster").Replace("410x231", "820x462"),
			Producer: Preset("Tokyo"),
			Sample:   Select("source").Attribute("src"),
			Series:   Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=genre\"]"),
			Release:  Match(`\d{4}/\d{2}/\d{2}`),
			Duration: Match(`\d{2}:\d{2}:\d{2}`),
			Id:       Select("input[name=\"product_uid\"]").Attribute("value"),
			Label:    Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=vendor\"], div.infowrapper > dl > dd > a[href^=\"/product/?vendor\"]"),
			Genre:    Select("div.infowrapper > dl > dd > a[href^=\"/product/?type=tag\"]"),
			Images:   Select("a[rel=\"cap\"]").Attribute("href").Replace(" ", "%20"),
		},
	}
}

func parseTokyoKey(key string) string {
	ext := path.Ext(key)
	name := strings.ToLower(strings.TrimSuffix(key, ext))
	re := regexp.MustCompile(`n-? ?\d{2,}`)

	matches := re.FindAllString(name, -1)

	if len(matches) > 0 {
		return matches[0]
	}
	return "nil"
}

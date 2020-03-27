package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
)

func Jav(id string) Site {
	dmmId := parseDmmKey(id)
	return Site{
		Key:       parseDmmKey(id),
		Url:       fmt.Sprintf("https://www.javbus.com/%s", dmmId),
		UserAgent: UserAgent,
		Path:      "dmm/$Actor/$Id $Title/",

		Selector: Selector{
			Title:    Select("h3"),
			Actor:    Select(".star-name"),
			Poster:   Select(".bigImage").Attribute("href"),
			Producer: Select("a[href^=\"https://www.javbus.com/studio/\"]"),
			Sample:   Select(".play-btn").Attribute("href"),
			Series:   Select("a[href^=\"https://www.javbus.com/series\"]"),
			Release:  Match(`\d{4}-\d{2}-\d{2}`),
			Duration: Match(`\d{0,4}分鐘`),
			Id:       Match(`識別碼: [A-Z]{0,6}-\d{0,6}`).Replace("識別碼: ", ""),
			Label:    Select("a[href^=\"https://www.javbus.com/label\"]"),
			Genre:    Select(".genre"),
			Images:   Select("a.sample-box").Attribute("href"),
		}.Extra("actors", Selects(".star-name")),
	}
}

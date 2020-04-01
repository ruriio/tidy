package sites

import (
	"fmt"
	. "github.com/ruriio/tidy/selector"
)

func JavLib(id string) Site {
	dmmId := parseDmmKey(id)
	url := fmt.Sprintf("http://www.javlibrary.com/cn/vl_searchbyid.php?keyword=%s", dmmId)
	return Site{
		Key:       parseDmmKey(id),
		Url:       url,
		UserAgent: UserAgent,
		//Cookies: []http.Cookie{{Name: "__cfduid", Value: "d1050bbdb76517956e3ea66542aa5b7c81584260542"},
		//	{Name: "cf_clearance", Value: "01026f8f689772cc62faa3993d7ccb7a49c805c5-1585746249-0-150"},
		//	{Name: "over18", Value: "18"},},

		Selector: Selector{
			Title:    Select("h3").Replace(dmmId, ""),
			Actor:    Select("span.cast"),
			Poster:   Select("#video_jacket_img").Attribute("src").Replace("//", "http://"),
			Producer: Select("#maker"),
			Sample:   Select(".play-btn").Attribute("href"),
			Series:   Select("a[href^=\"https://www.javbus.com/series\"]"),
			Release:  Match(`\d{4}-\d{2}-\d{2}`),
			Duration: Select("#video_length").Replace("长度:\n\t", ""),
			Id:       Select("div#video_id.item").Replace("识别码:\n\t", ""),
			Label:    Select("a[href^=\"https://www.javbus.com/label\"]"),
			Genre:    Select(".genre"),
			Images:   Select("a.sample-box").Attribute("href"),
		}.Extra("actors", Selects("span.cast")),
	}
}

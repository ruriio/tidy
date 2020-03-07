package main

import (
	"log"
	. "tidy/model"
	. "tidy/selector"
)

func Scrape(site Site) Meta {
	log.Printf("url: %s", site.Url)
	return site.Meta()
}

package main

import (
	. "tidy/model"
	. "tidy/selector"
)

func Scrape(site Site) Meta {
	return site.Meta()
}

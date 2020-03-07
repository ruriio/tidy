package main

import (
	"fmt"
	"testing"
	"tidy/sites"
)

func TestScrape(t *testing.T) {
	//meta := Scrape(sites.Dmm("ssni678"))
	meta := Scrape(sites.Fc2("1294320"))

	fmt.Println(meta.Json())
}

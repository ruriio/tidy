package main

import (
	"fmt"
	"testing"
	"tidy/sites"
)

func TestScrape(t *testing.T) {
	//meta := Scrape(sites.Dmm("ssni678"))
	//meta := Scrape(sites.Fc2("1294320"))
	meta := Scrape(sites.Fc2Club("437689"))

	fmt.Println(meta.Json())
}

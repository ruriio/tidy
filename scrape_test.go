package main

import (
	"fmt"
	"testing"
	"tidy/sites"
)

func TestScrape(t *testing.T) {
	//meta := Scrape(sites.Dmm("ssni678"))
	//meta := Scrape(sites.Fc2("1294320"))
	//meta := Scrape(sites.Fc2Club("437689"))
	//meta := Scrape(sites.Carib("030720-001"))
	//meta := Scrape(sites.CaribPr("022820_003"))
	//meta := Scrape(sites.Mgs("300MIUM-544"))
	//meta := Scrape(sites.Heyzo("2177"))
	//meta := Scrape(sites.Fantia("8209"))
	meta := Scrape(sites.Getchu("19622"))

	fmt.Println(meta.Json())
}

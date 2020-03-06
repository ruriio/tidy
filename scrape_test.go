package main

import (
	"fmt"
	"testing"
	"tidy/sites"
)

func TestScrape(t *testing.T) {
	meta := Scrape(sites.Dmm("ssni678"))

	fmt.Println(meta.Json())
}

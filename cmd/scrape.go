package cmd

import (
	. "fmt"
	. "github.com/ruriio/tidy/selector"
	. "github.com/ruriio/tidy/sites"
	. "github.com/ruriio/tidy/util"
	. "github.com/spf13/cobra"
	"log"
	"path"
	"path/filepath"
	"strings"
)

var scrapeCmd = &Command{
	Use:     "scrape",
	Aliases: []string{"dmm", "fc2", "mgs", "ave"},
	Short:   "Scrape site meta info",
	Long:    `Get site meta info`,
	Run:     run,
}

var siteMap = make(map[string]func(string) Site)

var extensions = map[string]bool{
	".mp4": true,
	".mkv": true,
	".wmv": true,
	".avi": true,
}

func run(cmd *Command, args []string) {

	initSites()

	if len(args) > 0 {
		siteId := cmd.CalledAs()
		id := args[0]
		if id == "." {
			scrapeDir(siteId)
		} else {
			scrape(siteId, id)
		}
	} else {
		Println("Need at least 1 args.")
	}
}

func initSites() {
	siteMap["dmm"] = Dmm
	siteMap["fc2"] = Fc2
	siteMap["mgs"] = Mgs
	siteMap["ave"] = Ave
}

func isSiteDir(name string) bool {
	return siteMap[name] != nil
}

func scrapeDir(siteId string) {
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		// ignore site dir
		if isSiteDir(file) {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file))
		if IsDirectory(file) || extensions[ext] {
			scrape(siteId, file)
		}
	}
}

func scrape(siteId string, id string) {
	site := siteMap[siteId](id)
	meta := site.Meta()

	if len(meta.Title) == 0 {
		return
	}

	dir := Sprintf("%v", meta.Extras["path"])
	dir = Move(id, dir)
	file := path.Join(dir, "meta.json")
	Write(file, meta.Byte())
	DownloadMedias(dir, meta.Poster, meta.Sample, meta.Images)
}

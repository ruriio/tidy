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
	"reflect"
	"runtime"
	"strings"
)

var siteMap = registerSites()

var scrapeCmd = &Command{
	Use:     "scrape",
	Aliases: getAliases(),
	Short:   "Scrape site meta info",
	Long:    `Get site meta info`,
	Run:     run,
}

var extensions = map[string]bool{
	".mp4": true,
	".mkv": true,
	".wmv": true,
	".avi": true,
}

func run(cmd *Command, args []string) {
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

func registerSites() map[string]func(string) Site {
	sites := make(map[string]func(string) Site)

	register(sites, Dmm)
	register(sites, Fc2)
	register(sites, Mgs)
	register(sites, Ave)
	register(sites, Tokyo)
	register(sites, Getchu)
	return sites
}

func register(sites map[string]func(string) Site, site func(string) Site) {
	key := strings.ToLower(getFuncName(site))
	sites[key] = site
}

func getAliases() []string {
	var aliases []string
	for site := range siteMap {
		aliases = append(aliases, site)
	}
	return aliases
}

func getFuncName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
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

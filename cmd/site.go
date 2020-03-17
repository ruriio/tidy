package cmd

import (
	. "fmt"
	. "github.com/spf13/cobra"
	"path"
	. "tidy/selector"
	. "tidy/sites"
	"tidy/util"
)

var siteCmd = &Command{
	Use:     "site",
	Aliases: []string{"dmm", "fc2"},
	Short:   "Site meta info",
	Long:    `Get site meta info`,
	Run: func(cmd *Command, args []string) {
		initSites()
		if len(args) > 0 {
			id := args[0]
			key := cmd.CalledAs()
			execute(id, key)

		} else {
			Println("Need at least 1 args.")
		}
	},
}

var siteMap = make(map[string]func(string) Site)

func initSites() {
	siteMap["dmm"] = Dmm
	siteMap["fc2"] = Fc2
}

func execute(id string, key string) {
	site := siteMap[key](id)
	meta := site.Meta()

	dir := meta.Extras["path"]
	file := path.Join(dir, "meta.json")
	util.Move(id, dir)
	util.Write(file, meta.Byte())
}

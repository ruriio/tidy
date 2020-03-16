package cmd

import (
	. "fmt"
	. "github.com/spf13/cobra"
	. "tidy/selector"
	. "tidy/sites"
)

var siteCmd = &Command{
	Use:     "site",
	Aliases: []string{"dmm", "fc2"},
	Short:   "Site meta info",
	Long:    `Get site meta info`,
	Run: func(cmd *Command, args []string) {
		initSites()
		if len(args) > 0 {
			site := siteMap[cmd.CalledAs()](args[0])
			Println(site.Meta().Json())
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

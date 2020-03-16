package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tidy",
	Long:  `All software has versions. This is Tidy's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tidy v1.0.0")
	},
}

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "tidy",
		Short: "Tidy is a helper to make your things tidy",
		Long:  `Tidy is a helper to make your things tidy.`,
	}
)

func Execute() error {
	addConfigs()
	addFlags()
	addCommands()

	return rootCmd.Execute()
}

func addConfigs() {
	cobra.OnInitialize(initConfigs)
}

func addCommands() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(siteCmd)
}

func addFlags() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"config file (default is $HOME/.tidy/config.yaml)")
}

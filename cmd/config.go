package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

func initConfigs() {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()

		if err != nil {
			er(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.tidy")
		viper.AddConfigPath("/etc/tidy/")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

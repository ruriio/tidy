package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	version := flag.Bool("version", false, "prints current version")
	v := flag.Bool("v", false, "prints current version")
	flag.Parse()

	if *version || *v {
		fmt.Println("Version: 1.0.0")
		os.Exit(0)
	}
}

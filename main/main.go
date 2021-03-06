package main

import (
	"paraglider/glider/glide"
	"flag"
	"os"
	"fmt"
	"strings"
)

func main() {

	// Default configuration file for paraglider
	yamlConfigFile := "/etc/paraglider/glider.yaml"

	// Optionally provide path of custom config file using command line argument
	flag.StringVar(&yamlConfigFile, "config", yamlConfigFile, "Path to configuration file")

	flag.Parse()
	args := flag.Args()

	if _, err := os.Stat(yamlConfigFile); os.IsNotExist(err) {
		fmt.Println("File does not exists")
		os.Exit(2)
	}

	if len(args) != 1 {
		fmt.Println("Invalid number of arguments provided")
		os.Exit(2)
	}

	action := strings.ToLower(args[0])

	if action != "start" && action != "stop" {
		fmt.Println("Invalid action provided. Action can be either of start|stop")
		os.Exit(2)
	}

	glide.Glide(yamlConfigFile, action)
}

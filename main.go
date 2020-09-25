package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arschles/aaronbot5000/pkg/cmd"
	"github.com/arschles/aaronbot5000/pkg/config"
)

var configFileName string

func main() {

	configFileName = os.Getenv("ERIKBOTDEV_CONFIG_FILE_NAME")
	if configFileName == "" {
		configFileName = "config.json"
	}
	log.Printf("Using config %s", configFileName)

	file, err := os.Open(configFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = config.LoadConfig(file)
	if err != nil {
		log.Fatalf("Error loading config (%s)", err)
	}

	cmd.Execute()
}

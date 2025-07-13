package main

import (
	"log"

	"github.com/mortal-coders/gitlab-mr-batch/internal/config"
)

func main() {

	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Config load error: %v", err)
	}

	log.Println(config)
}

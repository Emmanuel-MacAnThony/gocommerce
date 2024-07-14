package main

import (
	"log"

	"github.com/Emmanuel-MacAnThony/gocommerce/config"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/api"
)

func main() {
	cfg, err := config.SetUpEnv()

	if err != nil {
		log.Fatalf("Failed to load config")
	}

	api.StartServer(cfg)

	log.Println("Server Started 🚀🚀🚀")
}

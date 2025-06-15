package main

import (
	"log"

	"github.com/AlexeyTarasov77/messanger.chats/config"
	"github.com/AlexeyTarasov77/messanger.chats/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}

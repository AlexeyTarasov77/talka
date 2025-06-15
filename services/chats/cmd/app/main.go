package main

import (
	"github.com/AlexeyTarasov77/messanger.chats/config"
	"github.com/AlexeyTarasov77/messanger.chats/internal/app"
)

func main() {
	// Configuration
	cfg := config.MustLoad()

	// Run
	app.Run(cfg)
}

package main

import (
	"github.com/AlexeyTarasov77/messanger.users/config"
	"github.com/AlexeyTarasov77/messanger.users/internal/app"
)

func main() {
	// Configuration
	cfg := config.MustLoad()

	// Run
	app.Run(cfg)
}

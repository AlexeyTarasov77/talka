// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeyTarasov77/messanger.chats/config"
	"github.com/AlexeyTarasov77/messanger.chats/internal/controller/http"
	repo "github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage/postgres"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase/chats"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/httpserver"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	repositories := repo.NewRepositorories(pg)

	// Use-Case
	chatsUseCase := chats.New(repositories.Chats)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	http.NewRouter(httpServer.App, cfg, chatsUseCase, l)

	// Start servers
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

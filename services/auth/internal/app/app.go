// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeyTarasov77/messanger.users/config"
	"github.com/AlexeyTarasov77/messanger.users/internal/controller/http"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/security"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/sessions"
	repo "github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/mysql"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/redis"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase/auth"
	"github.com/AlexeyTarasov77/messanger.users/pkg/httpserver"
	"github.com/AlexeyTarasov77/messanger.users/pkg/jwt"
	"github.com/AlexeyTarasov77/messanger.users/pkg/logger"
	"github.com/AlexeyTarasov77/messanger.users/pkg/mysql"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	db, err := mysql.New(cfg.DB.URL, mysql.MaxPoolSize(cfg.DB.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.New: %w", err))
	}
	defer db.Close()
	rdb, err := redis_adapter.New(cfg.Redis.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis_adapter.New: %w", err))
	}
	defer rdb.Close(context.Background())
	repositories := repo.NewRepositorories(db)
	sessionManagerFactory := sessions.NewManagerFactory(rdb)
	jwtProvider, err := jwt.NewTokenProvider(cfg.Auth.TokenSecret, cfg.Auth.TokenAlg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - jwt.NewTokenProvider: %w", err))
	}
	// Use-Case
	authUseCase := auth.New(
		mysql.NewTransactionManager(db),
		repositories.Users,
		sessionManagerFactory,
		security.New(),
		jwtProvider,
		cfg.Auth.TokenTTL,
	)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	http.NewRouter(httpServer.App, cfg, authUseCase, l)

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

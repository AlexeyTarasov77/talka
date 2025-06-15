// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/AlexeyTarasov77/messanger.chats/config"
	_ "github.com/AlexeyTarasov77/messanger.chats/docs" // Swagger docs.
	v1 "github.com/AlexeyTarasov77/messanger.chats/internal/controller/http/v1"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *gin.Engine, cfg *config.Config, chatsUsecase usecase.Chats, log logger.Interface) {
	// Prometheus metrics
	if cfg.Metrics.Enabled {
		p := ginprometheus.NewPrometheus("messanger.chats")
		p.Use(app)
	}

	// Swagger
	// if cfg.Swagger.Enabled {
	// 	app.GET("/swagger/*", swagger.HandlerDefault)
	// }

	// K8s probe
	app.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.Writer.Flush()
	})

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewV1Routes(apiV1Group, chatsUsecase, log)
	}
}

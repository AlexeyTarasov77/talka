package v1

import (
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// NewTranslationRoutes -.
func NewV1Routes(v1Group *gin.RouterGroup, chatsUsecase usecase.Chats, log logger.Interface) {
	r := &V1{chatsUsecase: chatsUsecase, log: log, validator: validator.New(validator.WithRequiredStructEnabled())}

	chatsGroup := v1Group.Group("/chats")

	{
		chatsGroup.GET("/", r.listChats)
	}
}

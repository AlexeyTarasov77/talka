package v1

import (
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	chatsUsecase usecase.Chats
	log          logger.Interface
	validator    *validator.Validate
}

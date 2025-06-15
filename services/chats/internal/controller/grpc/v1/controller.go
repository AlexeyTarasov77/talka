package v1

import (
	v1 "github.com/AlexeyTarasov77/messanger.chats/docs/proto/v1"
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}

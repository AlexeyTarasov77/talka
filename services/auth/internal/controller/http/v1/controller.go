package v1

import (
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.users/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	authUsecase usecase.Auth
	log         logger.Interface
	validator   *validator.Validate
}

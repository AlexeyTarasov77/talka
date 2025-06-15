package v1

import (
	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/logger"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/rabbitmq/rmq_rpc/server"
	"github.com/go-playground/validator/v10"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(routes map[string]server.CallHandler, t usecase.Translation, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		routes["v1.getHistory"] = r.getHistory()
	}
}

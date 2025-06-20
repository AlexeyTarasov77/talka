package oauth_providers

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

type Interface interface {
	GetAuthURL(stateToken string) string
	GetAccessToken(ctx context.Context, authCode string) (string, error)
	FetchUserData(ctx context.Context) (*entity.User, error)
}

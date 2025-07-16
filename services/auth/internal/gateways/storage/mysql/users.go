package repo

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

type Users struct {
	Base
}

func (u *Users) GetByOAuthAccId(ctx context.Context, oauthAccID string) (*entity.User, error) {
	return nil, nil
}

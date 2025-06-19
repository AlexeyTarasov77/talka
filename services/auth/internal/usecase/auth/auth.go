package auth

import (
	"context"
	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
)

// UseCase -.
type UseCase struct {
	usersRepo gateways.UsersRepo
	txManager gateways.TransactionsManager
}

// New -.
func New(txManager gateways.TransactionsManager, usersRepo gateways.UsersRepo) *UseCase {
	return &UseCase{
		txManager: txManager,
		usersRepo: usersRepo,
	}
}

func (uc *UseCase) SignIn(ctx context.Context) (*entity.User, error)
func (uc *UseCase) SignUp(ctx context.Context) (*entity.User, error)

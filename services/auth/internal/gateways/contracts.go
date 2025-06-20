// Package gateways implements application outer layer logic. Each logic group in own file.
package gateways

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_gateways_test.go -package=usecase_test

type (
	UsersRepo interface {
		CheckExistsByIds(ctx context.Context, ids []int) (bool, error)
		Insert(ctx context.Context, user *entity.User) (*entity.User, error)
	}
	Transaction interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
	TransactionsManager interface {
		StartTransaction(ctx context.Context) (Transaction, error)
	}
	OAuthProvider interface {
		GetAuthURL(stateToken string) string
		GetAccessToken(ctx context.Context, authCode string) (string, error)
		FetchUserData(ctx context.Context) (*entity.User, error)
	}
	Storage interface {
		Get(key string) (string, error)
		Set(key, value string) error
	}
)

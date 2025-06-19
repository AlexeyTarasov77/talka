// Package gateways implements application outer layer logic. Each logic group in own file.
package gateways

import (
	"context"
	"time"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_gateways_test.go -package=usecase_test

type (
	ChatsRepo interface {
		GetAll(ctx context.Context) ([]entity.Chat, error)
		Save(ctx context.Context, chat entity.Chat) (entity.Chat, error)
		UpdateLastMsgInfo(ctx context.Context, chatId int, msgText string, msgDate time.Time) error
	}
	UsersRepo interface {
		CheckExistsByIds(ctx context.Context, ids []int) (bool, error)
	}
	Transaction interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
	TransactionsManager interface {
		StartTransaction(ctx context.Context) (Transaction, error)
	}
	SlugGenerator interface {
		GenerateRandomSlug() (string, error)
	}
)

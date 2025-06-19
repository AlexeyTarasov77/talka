// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/dto"
	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
)

type txCtxKeyType string

const TransactionCtxKey txCtxKeyType = "dbTrx"

func SetTransaction(ctx context.Context, tx any) context.Context {
	return context.WithValue(ctx, txCtxKeyType(TransactionCtxKey), tx)
}

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	Auth interface {
	}
)

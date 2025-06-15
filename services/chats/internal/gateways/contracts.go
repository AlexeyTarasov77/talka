// Package gateways implements application outer layer logic. Each logic group in own file.
package gateways

import "context"

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_gateways_test.go -package=usecase_test

type (
	ChatsRepo   interface{}
	Transaction interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
	TransactionsManager interface {
		StartTransaction(ctx context.Context) (Transaction, error)
	}
)

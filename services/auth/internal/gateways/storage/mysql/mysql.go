package repo

import (
	"context"
	"fmt"

	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage/mysql/sql/generated"
	"github.com/AlexeyTarasov77/messanger.users/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.users/pkg/mysql"
)

//go:generate mockgen -source=mysql.go -destination=mocks_repo_test.go -package=repo_test
//go:generate mockgen -source=../../../../pkg/mysql/mysql.go -destination=mocks_mysql_test.go -package=repo_test
type TxFactory interface {
	Begin() (mysql.QueryableTransaction, error)
}

type Base struct {
	DB TxFactory
}

func (r *Base) GetQueries(ctx context.Context) (*generated.Queries, error) {
	var dbTx mysql.QueryableTransaction
	tx := ctx.Value(usecase.TransactionCtxKey)
	if tx != nil {
		dbTx = tx.(generated.DBTX)
	} else {
		newTx, err := r.DB.Begin()
		if err != nil {
			return nil, fmt.Errorf("failed to begin new tx: %w", err)
		}
		dbTx = newTx
	}
	return generated.New(dbTx), nil
}

type Repositories struct {
	Users *Users
}

func NewRepositorories(m *mysql.MySQL) *Repositories {
	baseRepo := Base{m}
	return &Repositories{Users: &Users{baseRepo}}
}

package mysql

import (
	"context"
	"database/sql"

	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
)

type mySQLTransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *mySQLTransactionManager {
	return &mySQLTransactionManager{
		db: db,
	}
}

type mySQLTx struct {
	*sql.Tx
}

func (tx mySQLTx) Commit(ctx context.Context) error {
	commitErr := make(chan error, 1)
	go func() {
		commitErr <- tx.Tx.Commit()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-commitErr:
		return <-commitErr
	}
}

func (tx mySQLTx) Rollback(ctx context.Context) error {
	rollbackErr := make(chan error, 1)
	go func() {
		rollbackErr <- tx.Tx.Rollback()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-rollbackErr:
		return <-rollbackErr
	}
}

func (m *mySQLTransactionManager) StartTransaction(ctx context.Context) (gateways.Transaction, error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}
	return mySQLTx{tx}, nil
}

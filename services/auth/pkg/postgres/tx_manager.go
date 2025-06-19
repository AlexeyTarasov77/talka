package postgres

import (
	"context"

	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgTransactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) *pgTransactionManager {
	return &pgTransactionManager{
		pool: pool,
	}
}

func (m *pgTransactionManager) StartTransaction(ctx context.Context) (gateways.Transaction, error) {
	return m.pool.BeginTx(ctx, pgx.TxOptions{})
}

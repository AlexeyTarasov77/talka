package repo

import (
	"context"
	"fmt"

	"github.com/AlexeyTarasov77/messanger.chats/internal/gateways/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type queryBuilder interface {
	ToSql() (string, []any, error)
}

func ExecAndGetMany[T any](ctx context.Context, qb queryBuilder, pool *pgxpool.Pool, collectFunc pgx.RowToFunc[T]) ([]T, error) {
	if collectFunc == nil {
		collectFunc = pgx.RowToStructByNameLax[T]
	}
	queryable, err := GetQueryable(ctx, PgxPoolAdapter{pool})
	if err != nil {
		return nil, err
	}
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build sql query: %w", err)
	}
	rows, err := queryable.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute sql query: %w", err)
	}
	res, err := pgx.CollectRows(rows, collectFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to collect row into user struct: %w", err)
	}
	if len(res) == 0 {
		return nil, storage.ErrNotFound
	}
	return res, nil
}

func ExecAndGetOne[T any](ctx context.Context, qb queryBuilder, pool *pgxpool.Pool, collectFunc pgx.RowToFunc[T]) (*T, error) {
	res, err := ExecAndGetMany[T](ctx, qb, pool, collectFunc)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

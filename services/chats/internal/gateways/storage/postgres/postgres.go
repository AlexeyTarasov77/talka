//go:generate mockgen -source=postgres.go -destination=./mocks_postgres_test.go -package=repo_test

package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlexeyTarasov77/messanger.chats/internal/usecase"
	"github.com/AlexeyTarasov77/messanger.chats/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgErrCode string

const (
	UniqueViolationErrCode     pgErrCode = "23505"
	ForeignKeyViolationErrCode pgErrCode = "23503"
)

type Repositories struct {
	Chats *Chats
}

func NewRepositorories(pg *postgres.Postgres) *Repositories {
	return &Repositories{Chats: &Chats{pg}}
}

type Queryable interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type SupportsAcquire interface {
	Acquire(ctx context.Context) (Queryable, error)
}

func GetQueryable(ctx context.Context, connProvider SupportsAcquire) (Queryable, error) {
	tx := ctx.Value(usecase.TransactionCtxKey)
	if tx != nil {
		return tx.(Queryable), nil
	}
	conn, err := connProvider.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire conn from pool: %w", err)
	}
	return conn, nil
}

func getPgErrCode(err error) pgErrCode {
	var pgxErr *pgconn.PgError
	if errors.As(err, &pgxErr) {
		return pgErrCode(pgxErr.Code)
	}
	return ""
}

type PgxPoolAdapter struct {
	Pool *pgxpool.Pool
}

func (a PgxPoolAdapter) Acquire(ctx context.Context) (Queryable, error) {
	return a.Pool.Acquire(ctx)
}

package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/config"
	"github.com/jackc/pgx/v5"
)

func NewTestSuite(t *testing.T, ctx context.Context) (*Postgres, pgx.Tx) {
	os.Setenv("MODE", "test")
	cfg := config.MustLoad()
	pg, err := New(cfg.PG.URL)
	if err != nil {
		t.Fatal(err)
	}
	txManager := NewTransactionManager(pg.Pool)
	tx, err := txManager.StartTransaction(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil {
			t.Fatal(err)
		}
	})
	return pg, tx.(pgx.Tx)
}

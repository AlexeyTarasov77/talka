package mysql

import (
	"context"
	"os"
	"testing"

	"github.com/AlexeyTarasov77/messanger.users/config"
)

type MySQLTestSuite struct {
	t  *testing.T
	M  *MySQL
	Tx QueryableTransaction
}

func NewTestSuite(t *testing.T) *MySQLTestSuite {
	os.Setenv("MODE", "test")
	cfg := config.MustLoad()
	m, err := New(cfg.DB.URL)
	if err != nil {
		t.Fatal(err)
	}
	return &MySQLTestSuite{
		M: m,
		t: t,
	}
}

func (s *MySQLTestSuite) WithTx(ctx context.Context) *MySQLTestSuite {
	txManager := NewTransactionManager(s.M)
	tx, err := txManager.StartTransaction(ctx)
	if err != nil {
		s.t.Fatal(err)
	}
	s.t.Cleanup(func() {
		if err := tx.Rollback(ctx); err != nil {
			s.t.Fatal(err)
		}
	})

	return &MySQLTestSuite{
		M:  s.M,
		t:  s.t,
		Tx: tx.(mySQLTx).Tx,
	}
}

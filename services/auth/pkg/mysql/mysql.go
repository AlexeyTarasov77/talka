package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultMaxPoolSize = 1
	_defaultConnTimeout = time.Second
)

type MySQL struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	DB           *sql.DB
}

func New(url string, options ...Option) (*MySQL, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("failed opening mysql db connection: %w", err)
	}
	mysql := &MySQL{
		maxPoolSize: _defaultMaxPoolSize,
		connTimeout: _defaultConnTimeout,
		DB:          db,
	}
	for _, opt := range options {
		opt(mysql)
	}
	db.SetConnMaxLifetime(mysql.connTimeout)
	db.SetMaxOpenConns(mysql.maxPoolSize)
	db.SetMaxIdleConns(mysql.maxPoolSize)

	return mysql, nil
}

func (m *MySQL) Close() {
	if m.DB != nil {
		m.DB.Close()
	}
}

type QueryableTransaction interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func (m *MySQL) Begin() (QueryableTransaction, error) {
	return m.DB.Begin()
}

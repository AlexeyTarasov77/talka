package mysql

import (
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
	db           *sql.DB
}

func New(url string, options ...Option) (*MySQL, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("failed opening mysql db connection: %w", err)
	}
	mysql := &MySQL{
		maxPoolSize: _defaultMaxPoolSize,
		connTimeout: _defaultConnTimeout,
		db:          db,
	}
	for _, opt := range options {
		opt(mysql)
	}
	db.SetConnMaxLifetime(mysql.connTimeout)
	db.SetMaxOpenConns(mysql.maxPoolSize)
	db.SetMaxIdleConns(mysql.maxPoolSize)

	return mysql, nil
}

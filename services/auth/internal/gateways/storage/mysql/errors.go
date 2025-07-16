package repo

import (
	"database/sql"
	"errors"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/storage"
	mysqlerrs "github.com/go-sql-driver/mysql"
)

type mysqlErrCode int

const (
	UniqueViolationErrCode     mysqlErrCode = 1062
	ForeignKeyViolationErrCode mysqlErrCode = 1452
)

var errorsMapping = map[mysqlErrCode]error{
	UniqueViolationErrCode:     storage.ErrAlreadyExists,
	ForeignKeyViolationErrCode: storage.ErrRelationNotFound,
}

func MapErr(err error) error {
	if merr, ok := err.(*mysqlerrs.MySQLError); ok {
		return errorsMapping[mysqlErrCode(merr.Number)]
	}
	if errors.Is(err, sql.ErrNoRows) {
		return storage.ErrNotFound
	}
	return err
}

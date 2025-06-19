package repo

import (
	"github.com/AlexeyTarasov77/messanger.users/pkg/mysql"
)

type Users struct {
	db *mysql.MySQL
}

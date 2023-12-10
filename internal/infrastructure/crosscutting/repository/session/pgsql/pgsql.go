package pgsql

import (
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

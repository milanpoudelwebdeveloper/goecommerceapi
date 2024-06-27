package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewMyPostgresStorage(cfg string) (*sql.DB, error) {
	return sql.Open("postgres", cfg)

}

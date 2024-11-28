package storage

import (
	"database/sql"

	"github.com/gennadyterekhov/auth-microservice/internal/repositories"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func newDB(dsn string) (*sql.DB, error) {
	return sql.Open("pgx", dsn)
}

func NewRepo(dsn string) (*repositories.Repository, error) {
	db, err := newDB(dsn)
	if err != nil {
		return nil, err
	}
	return repositories.New(db), nil
}

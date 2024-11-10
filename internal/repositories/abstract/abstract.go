package abstract

import (
	"context"
	"database/sql"

	"github.com/gennadyterekhov/auth-microservice/internal/models"
)

const (
	ErrorNotFound                 = "not found"
	ErrorExpectedAtLeastOneResult = "sql: no rows in result set"
)

type QueryMaker interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Ping() error
	Close() error
}

// Scannable interface for generic scan behavior
type Scannable interface {
	ScanRow(row *sql.Row) error
	ScanRows(rows *sql.Rows) error
}

type Repository struct {
	DB QueryMaker
}

var _ Scannable = &models.User{}

func New(db QueryMaker) *Repository {
	return &Repository{
		DB: db,
	}
}

// Clear is used only in tests
func (repo *Repository) Clear() {
	queries := []string{
		"delete from users;",
	}

	for _, v := range queries {
		_, err := repo.DB.Exec(v)
		if err != nil {
			panic("error when clearing " + err.Error())
		}
	}
}

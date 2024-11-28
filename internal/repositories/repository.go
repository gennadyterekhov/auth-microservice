package repositories

import (
	"context"
	"database/sql"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
)

type queryMaker interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Ping() error
	Close() error
}

// Repository is a single point of access to the database. it's not divided between entities - one repo for all
type Repository struct {
	DB queryMaker
}

// Scannable interface for generic scan behavior
type Scannable interface {
	ScanRow(row *sql.Row) error
}

var _ Scannable = &models.User{}

func New(db queryMaker) *Repository {
	return &Repository{
		DB: db,
	}
}

var (
	_ interfaces.RepositoryInterface = NewErrorMock()
	_ interfaces.RepositoryInterface = New(nil)
)

// Clear is used only in tests
func (repo *Repository) Clear() {
	queries := []string{
		"delete from users;",
	}

	for _, v := range queries {
		_, err := repo.DB.Exec(v)
		if err != nil {
			logger.Errorln(err.Error())
		}
	}
}

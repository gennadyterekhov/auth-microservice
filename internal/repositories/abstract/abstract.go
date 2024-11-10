package abstract

import (
	"context"
	"database/sql"

	"github.com/gennadyterekhov/auth-microservice/internal/models"
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
}

var _ Scannable = &models.User{}

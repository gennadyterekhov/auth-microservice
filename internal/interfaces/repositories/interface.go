package repositories

import (
	"context"

	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
)

// RepositoryInterface contains methods required from 2 implementations - db and map(test). Naming is consistent with SQL
type RepositoryInterface interface {
	Clear()

	InsertUser(ctx context.Context, login string, password string) (*models2.User, error)
	SelectUserByID(ctx context.Context, id int64) (*models2.User, error)
	SelectUserByLogin(ctx context.Context, login string) (*models2.User, error)
}

package repositories

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/models"
)

type RepositoryInterface interface {
	Clear()

	InsertUser(ctx context.Context, login string, password string) (*models.User, error)
	SelectUserByID(ctx context.Context, id int64) (*models.User, error)
	SelectUserByLogin(ctx context.Context, login string) (*models.User, error)
}

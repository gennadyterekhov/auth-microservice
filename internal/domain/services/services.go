package services

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/categories"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/orders"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/users"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
)

type Services struct {
	Users      users.Service
	Orders     orders.Service
	Register   register.Service
	Login      auth.Service
	Categories categories.Service
}

func New(repo repositories.RepositoryInterface) *Services {
	return &Services{
		Categories: categories.NewService(repo),
		Users:      users.NewService(repo),
		Orders:     orders.NewService(repo),
		Register:   register.NewService(repo),
		Login:      auth.NewService(repo),
	}
}

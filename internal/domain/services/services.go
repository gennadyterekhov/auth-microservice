package services

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
)

type Services struct {
	Register *register.Service
	Login    *auth.Service
}

func New(repo repositories.RepositoryInterface) *Services {
	return &Services{
		Register: register.NewService(repo),
		Login:    auth.NewService(repo),
	}
}

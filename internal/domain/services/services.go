package services

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/login"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/register"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
)

// Services contains all so-called services (business logic irrespective of protocol)
type Services struct {
	Register *register.Service
	Login    *login.Service
}

func New(repo interfaces.RepositoryInterface) *Services {
	return &Services{
		Register: register.New(repo),
		Login:    login.New(repo),
	}
}

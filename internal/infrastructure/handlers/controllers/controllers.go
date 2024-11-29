package controllers

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/login"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/register"
)

type Controllers struct {
	Register *register.Controller
	Login    *login.Controller
}

func New(servs *services.Services) *Controllers {
	return &Controllers{
		Register: register.NewController(servs.Register),
		Login:    login.NewController(servs.Login),
	}
}

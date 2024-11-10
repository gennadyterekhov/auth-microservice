package controllers

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/login"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/register"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializer"
)

type Controllers struct {
	Register register.Controller
	Login    login.Controller
}

func NewControllers(servs *services.Services, pack *serializer.Pack) *Controllers {
	return &Controllers{
		Register: register.NewController(servs.Register, pack.Register),
		Login:    login.NewController(servs.Login, pack.Login),
	}
}

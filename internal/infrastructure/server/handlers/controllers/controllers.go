package controllers

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/health"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/login"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/register"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/serializers"
)

type Controllers struct {
	Register *register.Controller
	Login    *login.Controller
	Health   *health.Controller
}

func NewControllers(servs *services.Services, pack *serializers.Base) *Controllers {
	return &Controllers{
		Register: register.NewController(servs.Register, pack),
		Login:    login.NewController(servs.Login, pack),
		Health:   health.NewController(),
	}
}

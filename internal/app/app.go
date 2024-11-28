package app

import (
	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/controllers"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/swagger/swagrouter"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories"
	"github.com/gorilla/mux"
)

type App struct {
	addr     string
	router   *swagrouter.Router
	repo     *repositories.Repository
	services *services.Services
}

func New(dsn string) (*App, error) {
	app := &App{}

	repo, err := storage.NewRepo(dsn)
	if err != nil {
		return nil, err
	}

	servs := services.New(repo)
	serPack := serializers.New()
	controllersStruct := controllers.NewControllers(servs, serPack)

	routerInstance := swagrouter.NewRouter(controllersStruct)

	app.router = routerInstance
	app.repo = repo
	app.services = servs

	return app, nil
}

func (a App) Router() *mux.Router {
	return a.router.Router
}

func (a App) Repository() *repositories.Repository {
	return a.repo
}

func (a App) Services() *services.Services {
	return a.services
}

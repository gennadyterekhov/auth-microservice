package app

import (
	"fmt"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/controllers"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/swagger/swagrouter"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/project/config"
	"github.com/gennadyterekhov/auth-microservice/internal/repositories"
)

type App struct {
	ServerConfig *config.Config
	Repo         *repositories.Repository
	Router       *swagrouter.Router
}

func New() *App {
	app := &App{}

	serverConfig := config.New()

	repo := storage.NewRepo(serverConfig.DBDsn)

	servs := services.New(repo)
	serPack := serializers.New()
	controllersStruct := controllers.NewControllers(servs, serPack)

	routerInstance := swagrouter.NewRouter(controllersStruct)

	app.ServerConfig = serverConfig
	app.Repo = repo
	app.Router = routerInstance

	return app
}

func (a App) StartServer() error {
	fmt.Printf("Server started on %v\n", a.ServerConfig.Addr)
	err := http.ListenAndServe(a.ServerConfig.Addr, a.Router.Router)

	return err
}

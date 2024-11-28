package router

import (
	"net/http"
	"strings"

	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/controllers"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/health"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/middleware"
	"github.com/gorilla/mux"
)

type Router struct {
	Router      *mux.Router
	Controllers *controllers.Controllers
}

func NewRouter(controllers *controllers.Controllers) *Router {
	gorillaRouter := mux.NewRouter().StrictSlash(true)

	routerInstance := &Router{
		Router:      gorillaRouter,
		Controllers: controllers,
	}
	routerInstance.initializeRoutes()

	return routerInstance
}

func (r *Router) initializeRoutes() {
	type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
	type Routes []Route

	routes := Routes{
		Route{
			"AuthMicroserviceLogin",
			strings.ToUpper("Post"),
			"/login",
			http.HandlerFunc(r.Controllers.Login.Login),
		},

		Route{
			"AuthMicroserviceRegister",
			strings.ToUpper("Post"),
			"/register",
			http.HandlerFunc(r.Controllers.Register.Register),
		},

		Route{
			"Health",
			strings.ToUpper("Get"),
			"/health",
			http.HandlerFunc(health.Health),
		},
		// TODO add check, forget, verify-email endpoints
	}

	for _, route := range routes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(middleware.AddCommonMiddleware(route.HandlerFunc))
	}
}

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
			"options",
			"OPTIONS",
			"/",
			middleware.Logger(
				middleware.CorsAllowAll(
					http.HandlerFunc(health.Options),
				),
			).ServeHTTP,
		},

		Route{
			"Health",
			strings.ToUpper("Get"),
			"/health",
			middleware.Logger(
				middleware.CorsAllowAll(
					http.HandlerFunc(health.Health),
				),
			).ServeHTTP,
		},

		Route{
			"AuthMicroserviceLogin",
			strings.ToUpper("Post"),
			"/login",
			middleware.Logger(
				middleware.CorsAllowAll(
					middleware.ResponseContentTypeJSON(
						middleware.RequestContentTypeJSON(
							http.HandlerFunc(r.Controllers.Login.Login),
						),
					),
				),
			).ServeHTTP,
		},

		Route{
			"AuthMicroserviceRegister",
			strings.ToUpper("Post"),
			"/register",
			middleware.Logger(
				middleware.CorsAllowAll(
					middleware.ResponseContentTypeJSON(
						middleware.RequestContentTypeJSON(
							http.HandlerFunc(r.Controllers.Register.Register),
						),
					),
				),
			).ServeHTTP,
		},
		// TODO add forget, verify-email endpoints
	}

	for _, route := range routes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

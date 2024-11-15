package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func WithAuth(h http.Handler, middlewares ...Middleware) http.Handler {
	var allMiddlewares []Middleware
	allMiddlewares = append(allMiddlewares, Auth)
	allMiddlewares = append(allMiddlewares, middlewares...)

	return commonConveyor(h, allMiddlewares...)
}

func WithoutAuth(h http.Handler, middlewares ...Middleware) http.Handler {
	return commonConveyor(h, middlewares...)
}

func commonConveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	allMiddlewares := getCommonMiddlewares()
	allMiddlewares = append(allMiddlewares, middlewares...)
	return conveyor(h, allMiddlewares...)
}

func getCommonMiddlewares() []Middleware {
	return []Middleware{}
}

func conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
	middlewaresLength := len(middlewares)
	// in reverse, so that middlewares are applied in order that they are passed in router
	for i := middlewaresLength - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

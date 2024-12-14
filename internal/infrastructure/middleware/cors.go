package middleware

import (
	"net/http"
)

func CorsAllowAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(response, request)
	})
}

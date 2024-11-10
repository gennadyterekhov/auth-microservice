package middleware

import (
	"fmt"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

func RequestContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request == nil {
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		contentType := request.Header.Get("Content-Type")

		if contentType != "application/json" {
			endpoint := fmt.Sprintf("%v %v", request.Method, request.RequestURI)
			logger.Errorln(
				fmt.Sprintf(
					"RequestContentTypeJSON middleware (%v) failed, got %v",
					endpoint,
					contentType,
				),
			)
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(response, request)
	})
}

func ResponseContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request == nil {
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		response.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(response, request)
	})
}

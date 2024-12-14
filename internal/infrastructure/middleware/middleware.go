package middleware

import (
	"net/http"
)

func AddCommonMiddleware(handler http.Handler) http.Handler {
	handler = Logger(handler)
	handler = CorsAllowAll(handler)
	handler = RequestContentTypeJSON(handler)
	handler = ResponseContentTypeJSON(handler)
	return handler
}

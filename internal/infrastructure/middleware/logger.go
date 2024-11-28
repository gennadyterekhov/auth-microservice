package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, req)

		logger.Debugln(getMessage(req, start))
	})
}

func getMessage(req *http.Request, start time.Time) string {
	return fmt.Sprintf(
		`{"route":"%s %s", "body":%s, "time":"%s"}`,
		req.Method,
		req.RequestURI,
		getBodyAsString(req),
		time.Since(start),
	)
}

func getBodyAsString(req *http.Request) string {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return "error reading body"
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}

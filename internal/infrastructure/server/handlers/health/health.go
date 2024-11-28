package health

import (
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

func Health(res http.ResponseWriter, req *http.Request) {
	logger.Debugln("/api/health handler")

	_, err := res.Write([]byte(`{"status":"ok", "code":200}`))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

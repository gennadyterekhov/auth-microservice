package health

import (
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func Handler(controller *Controller) http.Handler {
	return http.HandlerFunc(controller.health)
}

func (controller *Controller) health(res http.ResponseWriter, req *http.Request) {
	logger.Debugln("/api/health handler")

	_, err := res.Write([]byte(`{"status":"ok", "code":200}`))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

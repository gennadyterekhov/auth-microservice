package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/login"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
)

type Controller struct {
	Service *login.Service
}

func NewController(service *login.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (controller *Controller) Login(res http.ResponseWriter, req *http.Request) {
	reqDto, err := getRequestDto(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	resDto, err := controller.Service.Login(req.Context(), reqDto)
	if err != nil {
		status := http.StatusInternalServerError

		if err.Error() == login.ErrorWrongCredentials {
			status = http.StatusUnauthorized
		}

		http.Error(res, err.Error(), status)
		return
	}
	res.Header().Set("Authorization", resDto.Token)

	err = serializers.WriteToWriter(res, resDto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func getRequestDto(req *http.Request) (*requests.Login, error) {
	requestDto := &requests.Login{
		Login:    "",
		Password: "",
	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(requestDto)
	if err != nil {
		return nil, err
	}

	if requestDto.Login == "" || requestDto.Password == "" {
		return nil, fmt.Errorf("login or password is empty")
	}

	return requestDto, nil
}

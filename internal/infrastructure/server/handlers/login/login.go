package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

type Controller struct {
	Service *auth.Service
}

func NewController(service *auth.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (controller *Controller) Login(res http.ResponseWriter, req *http.Request) {
	logger.Debugln("/api/user/login handler")

	reqDto, err := getRequestDto(req)
	if err != nil {
		logger.Errorln(err.Error())
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	resDto, err := controller.Service.Login(req.Context(), reqDto)
	if err != nil {
		status := http.StatusInternalServerError

		if err.Error() == auth.ErrorWrongCredentials {
			status = http.StatusUnauthorized
		}
		logger.Errorln(err.Error())

		http.Error(res, err.Error(), status)
		return
	}
	res.Header().Set("Authorization", resDto.Token)

	resBody, err := serializers.Serialize(resDto)
	if err != nil {
		logger.Errorln(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debugln("returning body", string(resBody))
	_, err = res.Write(resBody)
	if err != nil {
		logger.Errorln(err.Error())
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

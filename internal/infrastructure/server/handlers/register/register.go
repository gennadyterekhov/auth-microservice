package register

import (
	"encoding/json"
	"net/http"

	domain "github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

type Controller struct {
	Service *domain.Service
}

func NewController(service *domain.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (controller *Controller) Register(res http.ResponseWriter, req *http.Request) {
	logger.Debugln("/api/user/register handler")

	var err error
	reqDto, err := getRequestDto(req)
	if err != nil {
		logger.Errorln(err.Error())

		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if reqDto.Login == "" || reqDto.Password == "" {
		http.Error(res, "login or password is empty", http.StatusBadRequest)
	}

	resDto, err := controller.Service.Register(req.Context(), reqDto)
	if err != nil {
		status := http.StatusInternalServerError

		if err.Error() == domain.ErrorNotUniqueLogin {
			status = http.StatusConflict
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

func getRequestDto(req *http.Request) (*requests.Register, error) {
	requestDto := &requests.Register{
		Login:    "",
		Password: "",
	}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(requestDto)
	if err != nil {
		return nil, err
	}

	return requestDto, nil
}

package register

import (
	"encoding/json"
	"net/http"

	domain "github.com/gennadyterekhov/auth-microservice/internal/domain/register"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/handlers/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
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
	var err error
	reqDto, err := getRequestDto(req)
	if err != nil {
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

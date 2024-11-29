package factories

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/register"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/models/responses"
)

type Factory struct {
	repo interfaces.RepositoryInterface
}

func NewFactory(repo interfaces.RepositoryInterface) *Factory {
	return &Factory{
		repo: repo,
	}
}

func (f *Factory) RegisterForTest(login string, password string) *responses.Register {
	reqDto := &requests.Register{Login: login, Password: password}
	service := register.New(f.repo)
	resDto, err := service.Register(context.Background(), reqDto)
	if err != nil {
		panic(err)
	}

	return resDto
}

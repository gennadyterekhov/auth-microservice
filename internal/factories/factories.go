package factories

import (
	"context"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
)

type Factory struct {
	repo repositories.RepositoryInterface
}

func NewFactory(repo repositories.RepositoryInterface) *Factory {
	return &Factory{
		repo: repo,
	}
}

func (f *Factory) RegisterForTest(login string, password string) *responses.Register {
	reqDto := &requests.Register{Login: login, Password: password}
	service := register.NewService(f.repo)
	resDto, err := service.Register(context.Background(), reqDto)
	if err != nil {
		panic(err)
	}

	return resDto
}

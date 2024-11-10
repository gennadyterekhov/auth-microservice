package factories

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
)

type Factory struct {
	repo repositories.RepositoryInterface
}

func NewFactory(repo repositories.RepositoryInterface) *Factory {
	return &Factory{
		repo: repo,
	}
}

func (f *Factory) NewUser(name string) *models2.User {
	category, err := f.repo.InsertUser(context.Background(), name, "password", "")
	if err != nil {
		panic(err)
	}
	return category
}

func (f *Factory) RegisterForTest(login string, password string) *responses.Register {
	reqDto := &requests.Register{Login: login, Password: password}
	service := register.NewService(f.repo)
	resDto, err := service.Register(context.Background(), reqDto)
	if err != nil {
		panic(err)
	}

	err = auth.SetToken(resDto.Token)
	if err != nil {
		panic(err)
	}

	return resDto
}

func (f *Factory) RegisterRandForTest() *responses.Register {
	return f.RegisterForTest(fmt.Sprintf("user_%v", rand.Int()), "password")
}

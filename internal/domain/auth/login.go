package auth

import (
	"context"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/token"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/logger"
)

const ErrorWrongCredentials = "unknown credentials"

type Service struct {
	Repository repositories.RepositoryInterface
}

func NewService(repo repositories.RepositoryInterface) Service {
	return Service{
		Repository: repo,
	}
}

func (service *Service) Login(ctx context.Context, reqDto *requests.Login) (*responses.Login, error) {
	user, err := service.Repository.SelectUserByLogin(ctx, reqDto.Login)
	if err != nil {
		return nil, fmt.Errorf(ErrorWrongCredentials)
	}
	if user == nil {
		return nil, fmt.Errorf(ErrorWrongCredentials)
	}

	err = CheckPassword(reqDto.Password, user.Password)
	if err != nil {
		return nil, err
	}

	tokenString, err := token.CreateToken(user)
	if err != nil {
		return nil, err
	}

	resDto := responses.Login{
		Token: tokenString,
	}

	return &resDto, nil
}

func CheckPassword(plainPassword string, hashFromDB string) error {
	// ComparePasswordAndHash performs a constant-time comparison between a
	// plain-text password and Argon2id hash, using the parameters and salt
	// contained in the hash. It returns true if they match, otherwise it returns
	// false.
	match, err := argon2id.ComparePasswordAndHash(plainPassword, hashFromDB)
	if err != nil {
		logger.Errorln(err.Error())
		return err
	}

	if match {
		return nil
	}

	return fmt.Errorf(ErrorWrongCredentials)
}

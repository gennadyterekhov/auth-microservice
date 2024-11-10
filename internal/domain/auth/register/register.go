package register

import (
	"context"

	"github.com/alexedwards/argon2id"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/token"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
)

const ErrorNotUniqueLogin = "ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)"

type Service struct {
	Repository repositories.RepositoryInterface
}

func NewService(repo repositories.RepositoryInterface) Service {
	return Service{
		Repository: repo,
	}
}

func (service *Service) Register(ctx context.Context, reqDto *requests.Register) (*responses.Register, error) {
	encryptedPassword, err := encrypt(reqDto.Password)
	if err != nil {
		return nil, err
	}

	user, err := service.Repository.InsertUser(ctx, reqDto.Login, encryptedPassword, "")
	if err != nil {
		return nil, err
	}

	tokenString, err := token.CreateToken(user)
	if err != nil {
		return nil, err
	}

	resDto := responses.Register{
		ID:    user.ID,
		Token: tokenString,
	}

	return &resDto, nil
}

func encrypt(plainPassword string) (string, error) {
	// CreateHash returns an Argon2id hash of a plain-text password using the
	// provided algorithm parameters. The returned hash follows the format used
	// by the Argon2 reference C implementation and looks like this:
	// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG
	hash, err := argon2id.CreateHash(plainPassword, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, err
}

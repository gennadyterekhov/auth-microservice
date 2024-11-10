package repositories

import (
	"context"
	"fmt"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/stretchr/testify/mock"
)

// RepositoryErrorMock is used in tests in place of a real-DB repo. always returns error.
type RepositoryErrorMock struct {
	mock.Mock
	err error
}

var _ repositories.RepositoryInterface = &RepositoryErrorMock{}

func NewErrorMock() *RepositoryErrorMock {
	return &RepositoryErrorMock{
		err: fmt.Errorf("error"),
	}
}

func (repo *RepositoryErrorMock) Clear() {
}

// SetError can be used to customize error message or disabling it (passing nil)
func (repo *RepositoryErrorMock) SetError(err error) {
	repo.err = err
}

// users

func (repo *RepositoryErrorMock) InsertUser(context.Context, string, string) (*models2.User, error) {
	return nil, repo.err
}

func (repo *RepositoryErrorMock) SelectUserByID(context.Context, int64) (*models2.User, error) {
	return nil, repo.err
}

func (repo *RepositoryErrorMock) SelectUserByLogin(context.Context, string) (*models2.User, error) {
	return nil, repo.err
}

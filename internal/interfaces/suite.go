package interfaces

import (
	"bytes"
	"net/http/httptest"

	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/interfaces/repositories"
	"github.com/testcontainers/testcontainers-go"
)

type hasRepo interface {
	SetRepository(repo repositories.RepositoryInterface)
	GetRepository() repositories.RepositoryInterface
}

type hasDBContainer interface {
	SetDBContainer(cont testcontainers.Container)
	GetDBContainer() testcontainers.Container
}

type WithDb interface {
	hasRepo
	hasDBContainer
}

type WithFactory interface {
	WithDb
	SetFactory(fact factories.Interface)
	GetFactory() factories.Interface
}

type WithServer interface {
	WithFactory
	SetServer(srv *httptest.Server)
	GetServer() *httptest.Server

	SendGet(
		path string,
		token string,
	) (int, []byte)

	SendPostWithoutToken(
		path string,
		requestBody *bytes.Buffer,
	) int

	SendPost(
		path string,
		token string,
		requestBody *bytes.Buffer,
	) int

	SendPostAndReturnBody(
		path string,
		token string,
		requestBody *bytes.Buffer,
	) (int, []byte)
}

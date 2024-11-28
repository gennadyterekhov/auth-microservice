package interfaces

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/testcontainers/testcontainers-go"
)

type hasRepo interface {
	SetRepository(repo RepositoryInterface)
	GetRepository() RepositoryInterface
}

type hasDBContainer interface {
	SetDBContainer(cont testcontainers.Container)
	GetDBContainer() testcontainers.Container
}

type WithDb interface {
	T() *testing.T
	hasRepo
	hasDBContainer
}

type WithFactory interface {
	WithDb
	SetFactory(fact Interface)
	GetFactory() Interface
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

package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/suite"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type authTestSuite struct {
	suites.WithServer
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(authTestSuite))
}

func (suite *authTestSuite) SetupSuite() {
	inits.InitDbSuite(suite)
	inits.InitFactorySuite(suite)

	suite.SetServer(httptest.NewServer(getTestRouter()))
}

func getTestRouter() *chi.Mux {
	testRouter := chi.NewRouter()
	testRouter.Get(
		"/auth",
		WithAuth(
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(200)
			}),
		).ServeHTTP,
	)
	testRouter.Post(
		"/json",
		RequestContentTypeJSON(
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(200)
			}),
		).ServeHTTP,
	)
	testRouter.Post(
		"/luhn",
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(200)
		}).ServeHTTP,
	) //
	return testRouter
}

func (suite *authTestSuite) TestCanAuthWithToken() {
	resDto := suite.registerForTest("a", "a")

	responseStatusCode, _ := suite.SendGet(
		"/auth",
		resDto.Token,
	)

	assert.Equal(suite.T(), http.StatusOK, responseStatusCode)
}

func (suite *authTestSuite) Test401IfNoToken() {
	suite.registerForTest("a", "a")

	responseStatusCode, _ := suite.SendGet(
		"/auth",
		"incorrect token",
	)

	assert.Equal(suite.T(), http.StatusUnauthorized, responseStatusCode)
}

func (suite *authTestSuite) registerForTest(login string, password string) *responses.Register {
	reqDto := &requests.Register{Login: login, Password: password}
	service := register.NewService(suite.GetRepository())
	resDto, err := service.Register(context.Background(), reqDto)
	if err != nil {
		panic(err)
	}
	return resDto
}

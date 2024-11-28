package register

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/register"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/middleware"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type registerSuite struct {
	suites.WithServer
}

func TestRegisterHandler(t *testing.T) {
	suite.Run(t, new(registerSuite))
}

func (suite *registerSuite) SetupSuite() {
	inits.InitFactorySuite(suite)

	suite.SetServer(httptest.NewServer(
		middleware.AddCommonMiddleware(
			http.HandlerFunc(
				NewController(register.New(suite.GetRepository())).Register)),
	))
}

func (suite *registerSuite) TestCannotRegisterAlreadyPresent() {
	suite.GetFactory().RegisterForTest("a", "a")

	statusCode := suite.SendPost("/api/register", "", bytes.NewBufferString(`{"login":"a","password":"a"}`))
	require.NotEqual(suite.T(), 200, statusCode)
}

func (suite *registerSuite) TestCannotLoginWrongField() {
	statusCode := suite.SendPost("/api/register", "", bytes.NewBufferString(`{"loginn":"a","password":"a"}`))
	require.NotEqual(suite.T(), 200, statusCode)
}

func (suite *registerSuite) TestCanRegister() {
	statusCode := suite.SendPost("/api/register", "", bytes.NewBufferString(`{"login":"a","password":"a"}`))
	require.Equal(suite.T(), 200, statusCode)
}

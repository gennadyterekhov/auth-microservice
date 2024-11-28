package login

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type loginSuite struct {
	suites.WithServer
}

func TestLoginHandler(t *testing.T) {
	suite.Run(t, new(loginSuite))
}

func (suite *loginSuite) SetupSuite() {
	inits.InitFactorySuite(suite)

	suite.SetServer(httptest.NewServer(http.HandlerFunc(NewController(auth.NewService(suite.GetRepository())).Login)))
}

func (suite *loginSuite) TestCannotLoginNoUser() {
	statusCode := suite.SendPost("/api/login", "", bytes.NewBufferString(``))
	require.NotEqual(suite.T(), 200, statusCode)
}

func (suite *loginSuite) TestCannotLoginWrongField() {
	suite.GetFactory().RegisterForTest("a", "a")

	statusCode := suite.SendPost("/api/login", "", bytes.NewBufferString(`{"loginn":"a","password":"a"}`))
	require.NotEqual(suite.T(), 200, statusCode)
}

func (suite *loginSuite) TestCanLogin() {
	suite.GetFactory().RegisterForTest("a", "a")

	statusCode := suite.SendPost("/api/login", "", bytes.NewBufferString(`{"login":"a","password":"a"}`))
	require.Equal(suite.T(), 200, statusCode)
}

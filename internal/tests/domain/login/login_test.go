package login

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/login"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/token"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
	storageRepo "github.com/gennadyterekhov/auth-microservice/internal/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type loginTest struct {
	suites.WithFactory
	Service *login.Service
}

type loginErrorsTest struct {
	suites.WithFactory
	Service *login.Service
}

func TestLogin(t *testing.T) {
	suite.Run(t, new(loginTest))
}

func TestLoginErrors(t *testing.T) {
	suite.Run(t, new(loginErrorsTest))
}

func (suite *loginTest) SetupSuite() {
	inits.InitFactorySuite(suite)
	suite.Service = login.New(suite.GetRepository())
}

func (suite *loginErrorsTest) SetupSuite() {
	inits.InitFactorySuite(suite)
	suite.SetRepository(storageRepo.NewErrorMock())
	suite.Service = login.New(suite.GetRepository())
}

func (suite *loginTest) TestCanLogin() {
	factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	reqDto := &requests.Login{Login: "a", Password: "a"}
	resDto, err := suite.Service.Login(context.Background(), reqDto)
	assert.NoError(suite.T(), err)

	err = token.ValidateToken(resDto.Token, "a")
	assert.NoError(suite.T(), err)
}

func (suite *loginTest) TestCannotLoginWithWrongLogin() {
	factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	reqDto := &requests.Login{Login: "b", Password: "a"}
	_, err := suite.Service.Login(context.Background(), reqDto)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), login.ErrorWrongCredentials, err.Error())
}

func (suite *loginTest) TestCannotLoginWithWrongPassword() {
	factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	reqDto := &requests.Login{Login: "a", Password: "b"}
	_, err := suite.Service.Login(context.Background(), reqDto)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), login.ErrorWrongCredentials, err.Error())
}

func (suite *loginErrorsTest) TestLoginErrors() {
	reqDto := &requests.Login{Login: "a", Password: "a"}
	_, err := suite.Service.Login(context.Background(), reqDto)
	assert.Error(suite.T(), err)

	err = token.ValidateToken("", "a")
	assert.Error(suite.T(), err)

	err = login.CheckPassword("", "")
	assert.Error(suite.T(), err)
}

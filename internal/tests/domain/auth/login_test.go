package auth

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/token"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	storageRepo "github.com/gennadyterekhov/auth-microservice/internal/repositories"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type loginTest struct {
	suites.WithService
	Service *auth.Service
}

type loginErrorsTest struct {
	suites.WithService
	Service *auth.Service
}

func TestLogin(t *testing.T) {
	suite.Run(t, new(loginTest))
}

func TestLoginErrors(t *testing.T) {
	suite.Run(t, new(loginErrorsTest))
}

func (suite *loginTest) SetupSuite() {
	inits.InitServiceSuite(suite, nil)
	suite.Service = auth.NewService(suite.GetRepository())
}

func (suite *loginErrorsTest) SetupSuite() {
	inits.InitServiceSuite(suite, nil)
	suite.SetRepository(storageRepo.NewErrorMock())
	suite.Service = auth.NewService(suite.GetRepository())
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
	assert.Equal(suite.T(), auth.ErrorWrongCredentials, err.Error())
}

func (suite *loginTest) TestCannotLoginWithWrongPassword() {
	factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	reqDto := &requests.Login{Login: "a", Password: "b"}
	_, err := suite.Service.Login(context.Background(), reqDto)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), auth.ErrorWrongCredentials, err.Error())
}

func (suite *loginErrorsTest) TestLoginErrors() {
	reqDto := &requests.Login{Login: "a", Password: "a"}
	_, err := suite.Service.Login(context.Background(), reqDto)
	assert.Error(suite.T(), err)

	err = token.ValidateToken("", "a")
	assert.Error(suite.T(), err)

	err = auth.CheckPassword("", "")
	assert.Error(suite.T(), err)
}

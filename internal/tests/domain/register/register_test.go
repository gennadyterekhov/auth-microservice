package register

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/register"
	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suites.WithFactory
	Service *register.Service
}

func (suite *testSuite) SetupSuite() {
	inits.InitFactorySuite(suite)
	suite.Service = register.New(suite.GetRepository())
}

func TestRegisterDomain(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) TestCanRegister() {
	reqDto := &requests.Register{
		Login:    "a",
		Password: "a",
	}
	resDto, err := suite.Service.Register(context.Background(), reqDto)
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), "", resDto.Token)

	user, err := suite.Service.Repository.SelectUserByID(context.Background(), resDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "a", user.Login)
	assert.NotEqual(suite.T(), "a", user.Password)
}

func (suite *testSuite) TestCannotRegisterWhenLoginAlreadyUsed() {
	var err error
	_, err = suite.Service.Register(context.Background(), &requests.Register{Login: "a", Password: "a"})
	assert.NoError(suite.T(), err)
	_, err = suite.Service.Register(context.Background(), &requests.Register{Login: "a", Password: "a"})
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), register.ErrorNotUniqueLogin, err.Error())
}

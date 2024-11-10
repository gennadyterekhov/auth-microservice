package register

import (
	"context"
	"fmt"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/auth/register"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suites.WithService
	Service *register.Service
}

func (suite *testSuite) SetupSuite() {
	fmt.Println("(suite *testSuite) SetupSuite() in pkg")
	inits.InitServiceSuite(suite, nil)
	suite.Service = register.NewService(suite.GetRepository())
}

func TestRegisterDomain(t *testing.T) {
	fmt.Println("TestRegisterDomain")
	fmt.Println()

	inst := new(testSuite)
	fmt.Println("inst", inst)
	fmt.Printf("%+v \n\n", *inst)

	suite.Run(t, inst)
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
	assert.Equal(suite.T(), "ERROR: duplicate key value violates unique constraint \"users_login_key\" (SQLSTATE 23505)", err.Error())
}

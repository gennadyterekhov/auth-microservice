package categories

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/consts"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/categories"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suites.WithService
	Service categories.Service
}

func TestDomainCategories(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) SetupSuite() {
	inits.InitServiceSuite(suite, categories.NewService(suite.GetRepository()))
	suite.Service = categories.NewService(suite.GetRepository())
}

func (suite *testSuite) TestCanGetCategories() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	suite.createDifferentCategories()

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	all, err := suite.Service.GetAll(ctx)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 3, len(*all))
}

func (suite *testSuite) TestNoContentReturnsError() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	_, err := suite.Service.GetAll(ctx)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), categories.ErrorNoContent)
}

func (suite *testSuite) TestCanCreateCategory() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error

	_, err = suite.Service.Create(ctx, "test")
	assert.NoError(suite.T(), err)

	all, err := suite.Service.GetAll(ctx)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 1, len(*all))
	assert.Equal(suite.T(), "test", (*all)[0].Name)
}

func (suite *testSuite) createDifferentCategories() (*models.Category, *models.Category, *models.Category) {
	orderOldest, err := suite.Service.Repository.InsertCategory(
		context.Background(),
		"cat1",
	)
	assert.NoError(suite.T(), err)

	assert.NoError(suite.T(), err)
	orderMedium, err := suite.Service.Repository.InsertCategory(
		context.Background(),
		"cat2",
	)

	orderNewest, err := suite.Service.Repository.InsertCategory(
		context.Background(),
		"cat3",
	)
	assert.NoError(suite.T(), err)
	return orderNewest, orderMedium, orderOldest
}

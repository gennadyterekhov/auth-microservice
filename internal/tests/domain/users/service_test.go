package users

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/consts"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/users"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suites.WithService
	Service users.Service
}

type testErrorsSuite struct {
	suites.WithService
	Service users.Service
}

func newSuite() *testSuite {
	suiteInstance := new(testSuite)
	inits.InitServiceSuite(suiteInstance, nil)
	suiteInstance.Service = users.NewService(suiteInstance.GetRepository())

	return suiteInstance
}

func newSuiteWithMockRepo() *testErrorsSuite {
	suiteInstance := new(testErrorsSuite)
	inits.InitServiceSuite(suiteInstance, nil)
	suiteInstance.SetRepository(nil)
	suiteInstance.Service = users.NewService(suiteInstance.GetRepository())

	return suiteInstance
}

func TestUserService(t *testing.T) {
	suite.Run(t, newSuite())
}

func TestUserServiceErrors(t *testing.T) {
	suite.Run(t, newSuiteWithMockRepo())
}

func (suite *testSuite) TestCanUpdateBio() {
	suite.GetRepository().Clear()
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)

	var err error

	req := &users.UpdateBioRequest{"hello this is my new bio"}
	err = suite.Service.UpdateBio(ctx, req)
	assert.NoError(suite.T(), err)

	usr, err := suite.GetRepository().SelectUserByID(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "hello this is my new bio", usr.Bio)

	req = &users.UpdateBioRequest{"UPD newer bio"}
	err = suite.Service.UpdateBio(ctx, req)
	assert.NoError(suite.T(), err)

	usr, err = suite.GetRepository().SelectUserByID(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "UPD newer bio", usr.Bio)
}

func (suite *testSuite) TestCanUpdateCategoriesWhenEmptyRequestItDeletes() {
	suite.GetRepository().Clear()
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)

	var err error

	categoryFactory := factories.NewFactory(suite.GetRepository())
	cat3 := categoryFactory.NewCategory("cat3")
	req := &users.UpdateRequest{
		[]int64{cat3.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err := suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(cats))

	req = &users.UpdateRequest{[]int64{}}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err = suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, len(cats))
}

func (suite *testSuite) TestCanUpdateCategoriesWhenWasEmpty() {
	suite.GetRepository().Clear()
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)

	var err error
	cats, err := suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 0, len(cats))

	categoryFactory := factories.NewFactory(suite.GetRepository())
	cat1 := categoryFactory.NewCategory("cat1")
	cat2 := categoryFactory.NewCategory("cat2")
	categoryFactory.NewCategory("cat3")

	req := &users.UpdateRequest{
		[]int64{cat1.ID, cat2.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err = suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(cats))
}

func (suite *testSuite) TestCanUpdateCategoriesWhenAdding() {
	suite.GetRepository().Clear()
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)

	var err error

	categoryFactory := factories.NewFactory(suite.GetRepository())
	cat3 := categoryFactory.NewCategory("cat3")
	req := &users.UpdateRequest{
		[]int64{cat3.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err := suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(cats))

	cat1 := categoryFactory.NewCategory("cat1")
	cat2 := categoryFactory.NewCategory("cat2")

	req = &users.UpdateRequest{ // we must supply all 3 ids
		[]int64{cat1.ID, cat2.ID, cat3.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err = suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 3, len(cats))
}

func (suite *testSuite) TestCanUpdateCategoriesWhenOverwriting() {
	suite.GetRepository().Clear()
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error

	categoryFactory := factories.NewFactory(suite.GetRepository())
	cat3 := categoryFactory.NewCategory("cat3")
	req := &users.UpdateRequest{
		[]int64{cat3.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err := suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(cats))

	cat1 := categoryFactory.NewCategory("cat1")
	cat2 := categoryFactory.NewCategory("cat2")

	req = &users.UpdateRequest{
		[]int64{cat1.ID, cat2.ID},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.NoError(suite.T(), err)

	cats, err = suite.GetRepository().SelectUserCategories(ctx, userDto.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(cats))
}

func (suite *testErrorsSuite) TestErrorsUpdateBio() {
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, 1)

	var err error

	req := &users.UpdateBioRequest{"hello this is my new bio"}
	err = suite.Service.UpdateBio(ctx, req)
	assert.Error(suite.T(), err)
}

func (suite *testErrorsSuite) TestNeedContextUserToUpdate() {
	ctx := context.Background()
	var err error

	req := &users.UpdateRequest{
		[]int64{1},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.Error(suite.T(), err)
}

func (suite *testErrorsSuite) TestErrorsUpdateCategories() {
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, 1)
	var err error

	req := &users.UpdateRequest{
		[]int64{},
	}
	err = suite.Service.UpdateCategories(ctx, req)
	assert.Error(suite.T(), err)
}

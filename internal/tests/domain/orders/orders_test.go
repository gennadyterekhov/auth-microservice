package orders

import (
	"context"
	"testing"
	"time"

	"github.com/gennadyterekhov/auth-microservice/internal/consts"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/orders"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/requests"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	"github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suites.WithService
	Service orders.Service
}

func TestDomainOrders(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (suite *testSuite) SetupSuite() {
	inits.InitServiceSuite(suite, nil)
	suite.Service = orders.NewService(suite.GetRepository())
}

func (suite *testSuite) TestErrorIfNoUserInContext() {
	factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	var err error
	ctx := context.Background()
	_, err = suite.Service.GetAll(ctx)
	assert.Error(suite.T(), err)

	_, err = suite.Service.GetAllExceptAuthoredByCurrentUser(ctx)
	assert.Error(suite.T(), err)

	_, err = suite.Service.GetAllForUser(ctx)
	assert.Error(suite.T(), err)

	_, err = suite.Service.Create(ctx, &requests.Orders{})
	assert.Error(suite.T(), err)
}

func (suite *testSuite) TestCanSearchOrders() {
	var req *requests.SearchOrders
	var all []models.Order
	var err error

	ctxBack := context.Background()

	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	userDto2 := factories.NewFactory(suite.GetRepository()).RegisterForTest("b", "b")
	userDto3 := factories.NewFactory(suite.GetRepository()).RegisterForTest("c", "c")

	suite.Service.Repository.InsertOrder(ctxBack, userDto.ID, "1 hello title-substring", "hello desc-substr", 1)
	suite.Service.Repository.InsertOrder(ctxBack, userDto2.ID, "2 hello title-substring", "", 10)
	suite.Service.Repository.InsertOrder(ctxBack, userDto3.ID, "", "", 20)

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)

	req = &requests.SearchOrders{Title: "1 hello"}
	all, err = suite.Service.SearchOrders(ctx, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(all))

	req = &requests.SearchOrders{Title: "lo title-sub"}
	all, err = suite.Service.SearchOrders(ctx, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))

	req = &requests.SearchOrders{Title: "lo title-sub", Description: "desc"}
	all, err = suite.Service.SearchOrders(ctx, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(all))

	req = &requests.SearchOrders{PriceMin: 1, PriceMax: 10}
	all, err = suite.Service.SearchOrders(ctx, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))

	req = &requests.SearchOrders{PriceMin: 20, PriceMax: 20}
	all, err = suite.Service.SearchOrders(ctx, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(all))
}

func (suite *testSuite) TestCanGetOrders() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	orderNewest, orderMedium, orderOldest := suite.createDifferentOrders(userDto)

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	all, err := suite.Service.GetAll(ctx)
	assert.NoError(suite.T(), err)

	err = suite.Service.DeleteOrder(ctx, orderMedium.ID)
	assert.NoError(suite.T(), err)

	all, err = suite.Service.GetAll(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(all))
	assert.Equal(suite.T(), orderNewest.ID, (all)[0].ID)
	// assert.Equal(suite.T(), orderMedium.ID, (all)[1].ID)
	assert.Equal(suite.T(), orderOldest.ID, (all)[1].ID)
}

func (suite *testSuite) TestCanGetAllForUser() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	suite.createDifferentOrders(userDto)

	userDto2 := factories.NewFactory(suite.GetRepository()).RegisterForTest("b", "b")
	_, err := suite.Service.Repository.InsertOrder(context.Background(), userDto2.ID, "", "", 1)
	assert.NoError(suite.T(), err)

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	all, err := suite.Service.GetAllForUser(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 3, len(all))

	ctx = context.WithValue(context.Background(), consts.ContextUserIDKey, userDto2.ID)
	all, err = suite.Service.GetAllForUser(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(all))
}

func (suite *testSuite) TestCanGetAllExceptAuthoredByCurrentUser() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	suite.createDifferentOrders(userDto)

	userDto2 := factories.NewFactory(suite.GetRepository()).RegisterForTest("b", "b")
	_, err := suite.Service.Repository.InsertOrder(context.Background(), userDto2.ID, "", "", 1)
	assert.NoError(suite.T(), err)

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	all, err := suite.Service.GetAllExceptAuthoredByCurrentUser(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, len(all))

	ctx = context.WithValue(context.Background(), consts.ContextUserIDKey, userDto2.ID)
	all, err = suite.Service.GetAllExceptAuthoredByCurrentUser(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 3, len(all))
}

func (suite *testSuite) TestNoContentReturnsError() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	_, err := suite.Service.GetAll(ctx)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), err.Error(), orders.ErrorNoContent)
}

func (suite *testSuite) TestCanCreateOrder() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error

	cat1, err := suite.Service.Repository.InsertCategory(ctx, "cat1")
	cat2, err := suite.Service.Repository.InsertCategory(ctx, "cat2")

	req := &requests.Orders{"ttl", "desc", 1, []int64{cat1.ID, cat2.ID}}
	_, err = suite.Service.Create(ctx, req)
	assert.NoError(suite.T(), err)

	all, err := suite.Service.GetAll(ctx)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 1, len(all))
	assert.Equal(suite.T(), "ttl", (all)[0].Title)
	assert.Equal(suite.T(), "desc", (all)[0].Description)
}

func (suite *testSuite) TestCanBuyOrder() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	factories.NewFactory(suite.GetRepository()).RegisterForTest("buyer", "a")

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error
	buyer, err := suite.GetRepository().SelectUserByLogin(ctx, "buyer")

	req := &requests.Orders{"ttl", "desc", 1, []int64{}}
	order, err := suite.Service.Create(ctx, req)
	assert.NoError(suite.T(), err)

	err = suite.Service.Buy(ctx, order.ID, buyer.ID)
	assert.NoError(suite.T(), err)

	order, err = suite.Service.GetOrder(ctx, order.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), buyer.ID, *(order.BuyerID))
	assert.NotNil(suite.T(), order.ClosedAt)
}

func (suite *testSuite) TestCanUpdateOrder() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error

	order, err := suite.Service.Create(ctx, &requests.Orders{"ttl", "desc", 1, []int64{}})
	assert.NoError(suite.T(), err)

	cat1, err := suite.Service.Repository.InsertCategory(ctx, "cat1")
	cat2, err := suite.Service.Repository.InsertCategory(ctx, "cat2")

	req := &orders.UpdateRequest{
		Title:       "ttl 2",
		Description: "desc 2",
		Price:       2,
		CategoryIDs: []int64{cat1.ID, cat2.ID},
	}
	err = suite.Service.UpdateOrder(ctx, order.ID, req)
	assert.NoError(suite.T(), err)

	order, err = suite.Service.GetOrder(ctx, order.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "ttl 2", order.Title)
	assert.Equal(suite.T(), "desc 2", order.Description)
	assert.Equal(suite.T(), int64(2), order.Price)
}

func (suite *testSuite) TestCannotUpdateForeignOrder() {
	userDto1 := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	userDto2 := factories.NewFactory(suite.GetRepository()).RegisterForTest("b", "b")

	ctx1 := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto1.ID)
	ctx2 := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto2.ID)

	var err error

	order1, err := suite.Service.Create(ctx1, &requests.Orders{"ttl1", "desc", 1, []int64{}})
	assert.NoError(suite.T(), err)

	order2, err := suite.Service.Create(ctx2, &requests.Orders{"ttl2", "desc", 1, []int64{}})
	assert.NoError(suite.T(), err)

	req := &orders.UpdateRequest{Title: "ttl 2", Description: "desc 2", Price: 2, CategoryIDs: []int64{}}
	err = suite.Service.UpdateOrder(ctx1, order2.ID, req)
	assert.Error(suite.T(), err)

	err = suite.Service.UpdateOrder(ctx2, order1.ID, req)
	assert.Error(suite.T(), err)

	// cannot update when no user at all
	err = suite.Service.UpdateOrder(context.Background(), order1.ID, req)
	assert.Error(suite.T(), err)
}

func (suite *testSuite) TestCanDeleteOrder() {
	userDto := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")

	ctx := context.WithValue(context.Background(), consts.ContextUserIDKey, userDto.ID)
	var err error

	order, err := suite.Service.Create(ctx, &requests.Orders{"ttl", "desc", 1, []int64{}})
	assert.NoError(suite.T(), err)

	err = suite.Service.DeleteOrder(ctx, order.ID)
	assert.NoError(suite.T(), err)

	order, err = suite.Service.GetOrder(ctx, order.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), order)
	assert.NotNil(suite.T(), order.DeletedAt)
}

func (suite *testSuite) createDifferentOrders(
	userDto *responses.Register,
) (*models.Order, *models.Order, *models.Order) {
	orderOldest, err := suite.Service.Repository.InsertOrder(
		context.Background(),
		userDto.ID,
		//&ten,
		"", "",
		1,
	)
	assert.NoError(suite.T(), err)

	time.Sleep(time.Millisecond * 1)

	assert.NoError(suite.T(), err)
	orderMedium, err := suite.Service.Repository.InsertOrder(
		context.Background(),
		userDto.ID,
		//&ten,
		"", "",
		1,
	)

	time.Sleep(time.Millisecond * 1)

	orderNewest, err := suite.Service.Repository.InsertOrder(
		context.Background(),
		userDto.ID,
		//&ten,
		"", "",
		1,
	)
	assert.NoError(suite.T(), err)
	return orderNewest, orderMedium, orderOldest
}

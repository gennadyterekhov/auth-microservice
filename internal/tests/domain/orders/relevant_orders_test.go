package orders

import (
	"context"
	"time"

	"github.com/gennadyterekhov/auth-microservice/internal/consts"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/orders"
	"github.com/gennadyterekhov/auth-microservice/internal/domain/users"
	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/gennadyterekhov/auth-microservice/internal/factories"
	models2 "github.com/gennadyterekhov/auth-microservice/internal/models"
	"github.com/stretchr/testify/assert"
)

func (suite *testSuite) TestCanGetRelevant() {
	usrSrv := users.NewService(suite.GetRepository())
	var all []models2.Order
	var err error

	userWhoSearched := factories.NewFactory(suite.GetRepository()).RegisterForTest("a", "a")
	userWhoCreatedOrders := factories.NewFactory(suite.GetRepository()).RegisterForTest("b", "b")

	ctxBack := context.Background()
	ctx := context.WithValue(ctxBack, consts.ContextUserIDKey, userWhoCreatedOrders.ID)

	catPhoto, catModel, catMusic, catIT := suite.prepareCategories(ctx, err)

	suite.prepareData(ctx, userWhoCreatedOrders, catPhoto, catModel, catMusic, catIT)

	ctx = context.WithValue(ctxBack, consts.ContextUserIDKey, userWhoSearched.ID)
	usrSrv.UpdateCategories(ctx, &users.UpdateRequest{[]int64{catPhoto.ID, catModel.ID}})

	all, err = suite.Service.GetRelevant(ctx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 3, len(all))
	assert.Equal(suite.T(), "only model", all[0].Title)
	assert.Equal(suite.T(), "only photo", all[1].Title)
	assert.Equal(suite.T(), "all cats", all[2].Title)
}

func (suite *testSuite) prepareData(ctx context.Context, author *responses.Register, catPhoto, catModel, catMusic, catIT *models2.Category) (*models2.Order, *models2.Order, *models2.Order, *models2.Order, *models2.Order) {
	var err error
	var req *orders.UpdateRequest

	orderEmpty, err := suite.Service.Repository.InsertOrder(ctx, author.ID, "empty", "", 1)
	time.Sleep(time.Millisecond)
	orderAllCats, err := suite.Service.Repository.InsertOrder(ctx, author.ID, "all cats", "", 1)
	time.Sleep(time.Millisecond)
	orderPhoto, err := suite.Service.Repository.InsertOrder(ctx, author.ID, "only photo", "", 1)
	time.Sleep(time.Millisecond)
	orderMusicIT, err := suite.Service.Repository.InsertOrder(ctx, author.ID, "music it", "", 1)
	time.Sleep(time.Millisecond)
	orderModel, err := suite.Service.Repository.InsertOrder(ctx, author.ID, "only model", "", 1)

	req = &orders.UpdateRequest{Title: orderAllCats.Title, CategoryIDs: []int64{catPhoto.ID, catModel.ID, catMusic.ID, catIT.ID}}
	err = suite.Service.UpdateOrder(ctx, orderAllCats.ID, req)

	req = &orders.UpdateRequest{Title: orderPhoto.Title, CategoryIDs: []int64{catPhoto.ID}}
	err = suite.Service.UpdateOrder(ctx, orderPhoto.ID, req)

	req = &orders.UpdateRequest{Title: orderMusicIT.Title, CategoryIDs: []int64{catIT.ID, catMusic.ID}}
	err = suite.Service.UpdateOrder(ctx, orderMusicIT.ID, req)

	req = &orders.UpdateRequest{Title: orderModel.Title, CategoryIDs: []int64{catModel.ID}}
	err = suite.Service.UpdateOrder(ctx, orderModel.ID, req)

	assert.NoError(suite.T(), err)
	return orderEmpty, orderAllCats, orderPhoto, orderMusicIT, orderModel
}

func (suite *testSuite) prepareCategories(ctx context.Context, err error) (*models2.Category, *models2.Category, *models2.Category, *models2.Category) {
	catPhoto, err := suite.Service.Repository.InsertCategory(ctx, "photo")
	catModel, err := suite.Service.Repository.InsertCategory(ctx, "model")
	catMusic, err := suite.Service.Repository.InsertCategory(ctx, "music")
	catIT, err := suite.Service.Repository.InsertCategory(ctx, "IT")
	return catPhoto, catModel, catMusic, catIT
}

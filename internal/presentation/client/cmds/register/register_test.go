package register

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/client/swagger/models"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/handlers/controllers"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/server/swagger/swagrouter"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/serializers"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type registerCmdSuite struct {
	suites.WithServer
}

func TestRegisterCmd(t *testing.T) {
	suite.Run(t, new(registerCmdSuite))
}

func (suite *registerCmdSuite) SetupSuite() {
	inits.InitFactorySuite(suite)
	servs := services.New(suite.GetRepository())
	controllersStruct := controllers.NewControllers(servs, serializers.New())

	suite.SetServer(httptest.NewServer(swagrouter.NewRouter(controllersStruct).Router))
}

func (suite *registerCmdSuite) TestCanRegisterNoCommand() {
	ctx := context.Background()
	conf := swagger.NewConfiguration(false)
	conf.BasePath = suite.GetServer().URL
	client := swagger.NewAPIClient(conf)

	grpcResp, _, err := client.AuthApi.ArtDealersRegister(ctx, models.ProtobufRegisterRequest{Login: "a", Password: "a"})
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "", grpcResp.Error_)
	assert.NotEqual(suite.T(), "", grpcResp.Token)
}

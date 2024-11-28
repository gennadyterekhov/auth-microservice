package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type healthSuite struct {
	suites.WithServer
}

func TestHealthHandler(t *testing.T) {
	suite.Run(t, new(healthSuite))
}

func (suite *healthSuite) SetupSuite() {
	inits.InitFactorySuite(suite)

	suite.SetServer(httptest.NewServer(http.HandlerFunc(Health)))
}

func (suite *healthSuite) TestHealth() {
	statusCode, _ := suite.SendGet("/health", "")
	require.Equal(suite.T(), 200, statusCode)
}

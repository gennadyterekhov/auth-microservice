package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/tests/inits"
	"github.com/gennadyterekhov/auth-microservice/internal/tests/suites"
	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type jsonTestSuite struct{ suites.WithServer }

func TestJSON(t *testing.T) {
	suite.Run(t, new(jsonTestSuite))
}

func (suite *jsonTestSuite) SetupSuite() {
	inits.InitDbSuite(suite)
	inits.InitFactorySuite(suite)

	suite.SetServer(httptest.NewServer(getTestRouter()))
}

func (suite *jsonTestSuite) TestCanSendIfJson() {
	path := "/json"
	req, err := http.NewRequest(http.MethodPost, suite.GetServer().URL+path, nil)
	require.NoError(suite.T(), err)
	req.Header.Set("Content-Type", "application/json")

	response, err := suite.GetServer().Client().Do(req)
	require.NoError(suite.T(), err)
	response.Body.Close()

	assert.Equal(suite.T(), http.StatusOK, response.StatusCode)
}

func (suite *jsonTestSuite) Test400IfNotJson() {
	path := "/json"
	req, err := http.NewRequest(http.MethodPost, suite.GetServer().URL+path, nil)
	require.NoError(suite.T(), err)

	response, err := suite.GetServer().Client().Do(req)
	require.NoError(suite.T(), err)
	response.Body.Close()

	assert.Equal(suite.T(), http.StatusBadRequest, response.StatusCode)
}

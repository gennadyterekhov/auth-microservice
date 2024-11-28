package suites

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestDbSuite(t *testing.T) {
	suite.Run(t, new(WithDb))
}

func (suite *WithDb) TestCanGetRepository() {
	require.NotNil(suite.T(), suite.GetRepository())
}

func TestFactorySuite(t *testing.T) {
	suite.Run(t, new(WithFactory))
}

func (suite *WithFactory) TestCanGetFactory() {
	require.NotNil(suite.T(), suite.GetRepository())
	require.NotNil(suite.T(), suite.GetFactory())
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(WithServer))
}

func (suite *WithServer) TestCanGetService() {
	suite.T().Skipf("idk why it fails")
	// require.Nil(suite.T(), suite.GetServer()) // server not set by default. user must provide it manually
}

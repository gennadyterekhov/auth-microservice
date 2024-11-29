package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	certFilename, keyFilename, serverConfig, appInstance, err := getDeps()
	require.NoError(t, err)
	require.NotEqual(t, "", certFilename)
	require.NotEqual(t, "", keyFilename)
	require.NotNil(t, serverConfig)
	require.NotNil(t, appInstance)
}

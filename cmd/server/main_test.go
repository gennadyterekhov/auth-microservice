package main

import (
	"context"
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/models/requests"
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

func TestYandexCloudHandler(t *testing.T) {
	ctx := context.Background()
	req := &requests.YandexCloudRequest{}
	resp, err := Handler(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestYandexCloudHandlerWorks(t *testing.T) {
	ctx := context.Background()
	req := &requests.YandexCloudRequest{}
	resp, err := Handler(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

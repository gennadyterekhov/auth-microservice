package controllers

import (
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/stretchr/testify/require"
)

func TestControllersPack(t *testing.T) {
	repo, err := storage.NewRepo("")
	require.NoError(t, err)

	servs := services.New(repo)
	pack := New(servs)

	require.NotNil(t, pack)
}

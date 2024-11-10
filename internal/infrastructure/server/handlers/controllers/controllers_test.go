package controllers

import (
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/domain/services"
	"github.com/gennadyterekhov/auth-microservice/internal/infrastructure/storage"
	"github.com/gennadyterekhov/auth-microservice/internal/presentation/serializers"
	"github.com/stretchr/testify/require"
)

func TestControllersPack(t *testing.T) {
	servs := services.New(storage.NewRepo(""))
	serPack := serializers.New()
	pack := NewControllers(servs, serPack)

	require.NotNil(t, pack)
}

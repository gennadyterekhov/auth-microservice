package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrations(t *testing.T) {
	ctx := context.Background()
	container, _, err := CreatePostgresContainerAndRunMigrations(ctx)
	require.NoError(t, err)

	require.NoError(t, container.Terminate(ctx))
}

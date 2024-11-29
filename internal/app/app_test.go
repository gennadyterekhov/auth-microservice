package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApp(t *testing.T) {
	// no error with empty dsn
	inst, err := New("")
	require.NoError(t, err)
	require.NotNil(t, inst)

	require.NotNil(t, inst.Router())
	require.NotNil(t, inst.Repository())
	require.NotNil(t, inst.Services())
}

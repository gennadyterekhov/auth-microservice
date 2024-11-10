package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanGetToken(t *testing.T) {
	var err error

	realToken, err := GetToken()
	require.NoError(t, err)

	err = SetToken("hello")
	require.NoError(t, err)

	tok, err := GetToken()

	require.NoError(t, err)
	require.Equal(t, "hello", tok)

	err = SetToken(realToken)
	require.NoError(t, err)
}

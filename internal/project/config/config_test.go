package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	conf, err := New()
	require.NoError(t, err)
	require.Equal(t, defaultAddr, conf.Addr)
	require.Equal(t, defaultDbUrl, conf.DBDsn)
}

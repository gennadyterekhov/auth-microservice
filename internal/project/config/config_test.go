package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	conf := New()

	require.Equal(t, defaultAddr, conf.Addr)
	require.Equal(t, defaultDbUrl, conf.DBDsn)
}

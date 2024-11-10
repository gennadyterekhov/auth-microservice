package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfigValues(t *testing.T) {
	conf := New()

	assert.Equal(t, "host=localhost user=authmcrsrv_user password=authmcrsrv_pass dbname=authmcrsrv_db sslmode=disable", conf.DBDsn)
	assert.Equal(t, "localhost:8080", conf.Addr)
}

func TestEnvVarsOverwriteCliFlags(t *testing.T) {
	cmd := os.Args[0]
	os.Args = []string{cmd, "-d=2"}

	err := os.Setenv("DATABASE_URI", "1")
	assert.NoError(t, err)

	conf := New()
	assert.Equal(t, "1", conf.DBDsn)
}

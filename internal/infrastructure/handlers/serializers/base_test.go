package serializers

import (
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/models/responses"
	"github.com/stretchr/testify/require"
)

func TestCanSerialize(t *testing.T) {
	log := &responses.Login{Token: "hello"}
	bts, err := serialize(log)

	require.NoError(t, err)
	require.Equal(t, `{"token":"hello"}`, string(bts))

	reg := &responses.Register{ID: 1, Token: "hello"}
	bts, err = serialize(reg)

	require.NoError(t, err)
	require.Equal(t, `{"id":1,"token":"hello"}`, string(bts))
}

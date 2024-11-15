package serializers

import (
	"testing"

	"github.com/gennadyterekhov/auth-microservice/internal/dtos/responses"
	"github.com/stretchr/testify/require"
)

func TestCanSerialize(t *testing.T) {
	s := New()

	log := &responses.Login{Token: "hello"}
	bts, err := s.Serialize(log)

	require.NoError(t, err)
	require.Equal(t, `{"token":"hello"}`, string(bts))

	reg := &responses.Register{ID: 1, Token: "hello"}
	bts, err = s.Serialize(reg)

	require.NoError(t, err)
	require.Equal(t, `{"id":1,"token":"hello"}`, string(bts))
}

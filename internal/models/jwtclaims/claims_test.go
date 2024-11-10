package jwtclaims

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClaims(t *testing.T) {
	claims := &Claims{}

	var err error

	_, err = claims.GetExpirationTime()
	assert.NoError(t, err)

	_, err = claims.GetIssuedAt()
	assert.NoError(t, err)

	_, err = claims.GetNotBefore()
	assert.NoError(t, err)

	_, err = claims.GetIssuer()
	assert.NoError(t, err)

	_, err = claims.GetSubject()
	assert.NoError(t, err)

	_, err = claims.GetAudience()
	assert.NoError(t, err)

	_, err = claims.GetUserID()
	assert.NoError(t, err)
}

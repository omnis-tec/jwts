package tests

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJwkGetSet(t *testing.T) {
	jwks := app.core.Jwk.GetSet()
	require.NotNil(t, jwks)

	jwksBytes, err := json.MarshalIndent(jwks, "", "  ")
	require.Nil(t, err)
	require.Greater(t, len(jwksBytes), 0)
}

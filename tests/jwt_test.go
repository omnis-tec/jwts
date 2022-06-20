package tests

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJwtCreate(t *testing.T) {
	// "123", 86400, map[string]any{"hello": "world", "usr_type": "77"}
	jwt, err := app.core.Jwt.Create(map[string]any{
		"sub":         "123",
		"exp_seconds": 86400,
		"hello":       "world",
		"usr_type":    "77",
	})
	require.Nil(t, err)

	validateRepObj, err := app.core.Jwt.Validate(jwt.Token)
	require.Nil(t, err)
	require.True(t, validateRepObj.Valid)
	require.Equal(t, "world", validateRepObj.Claims["hello"])
	require.Equal(t, "77", validateRepObj.Claims["usr_type"])
}

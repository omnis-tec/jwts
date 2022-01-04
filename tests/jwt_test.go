package tests

import (
	_ "embed"
	"testing"

	"github.com/mechta-market/jwts/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestJwtCreate(t *testing.T) {
	// "123", 86400, map[string]interface{}{"hello": "world", "usr_type": "77"}
	jwt, err := app.ucs.JwtCreate(&entities.JwtCreateReqSt{
		Sub:        "123",
		ExpSeconds: 86400,
		Payload:    map[string]interface{}{"hello": "world", "usr_type": "77"},
	})
	require.Nil(t, err)

	validateRepObj, err := app.ucs.JwtValidate(jwt)
	require.Nil(t, err)
	require.True(t, validateRepObj.Valid)
	require.Equal(t, "world", validateRepObj.Claims["hello"])
	require.Equal(t, "77", validateRepObj.Claims["usr_type"])
}

package core

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mechta-market/jwts/internal/cns"
	"github.com/mechta-market/jwts/internal/domain/entities"
	"github.com/mechta-market/jwts/internal/domain/errs"
)

type Jwt struct {
	r *St
}

func NewJwt(r *St) *Jwt {
	return &Jwt{
		r: r,
	}
}

func (c *Jwt) Create(reqClaims map[string]interface{}) (string, error) {
	if c.r.privateKey == nil {
		return "", nil
	}

	claims := jwt.MapClaims{}

	for k, v := range reqClaims {
		claims[k] = v
	}

	now := time.Now()

	claims["iss"] = cns.JwtIssuer
	if expSeconds, ok := claims["exp_seconds"]; ok {
		expSecondsStr := fmt.Sprintf("%v", expSeconds)
		if expSecondsInt, err := strconv.ParseInt(expSecondsStr, 10, 64); err == nil {
			claims["exp"] = now.Unix() + expSecondsInt
			delete(claims, "exp_seconds")
		}
	}
	claims["iat"] = now.Add(-5 * time.Second).Unix()

	t := jwt.NewWithClaims(jwt.GetSigningMethod(cns.JwtSigningMethod), claims)

	if c.r.kid != "" {
		t.Header["kid"] = c.r.kid
	}

	return t.SignedString(c.r.privateKey)
}

func (c *Jwt) Validate(value string) (*entities.JwtValidateRepSt, error) {
	result := &entities.JwtValidateRepSt{}

	if c.r.publicKey == nil {
		return nil, errs.ServiceNA
	}

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(value, &claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errs.InvalidToken
		}
		return c.r.publicKey, nil
	})
	if err == nil {
		result.Valid = true
	}

	result.Claims = claims

	return result, nil
}

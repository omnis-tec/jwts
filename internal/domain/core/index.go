package core

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg logger.Lite

	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	kid        string

	Jwk *Jwk
	Jwt *Jwt
}

func New(
	lg logger.Lite,
) *St {
	c := &St{
		lg: lg,
	}

	c.Jwk = NewJwk(c)
	c.Jwt = NewJwt(c)

	return c
}

func (c *St) SetKeys(privateKeyBytes []byte, publicKeyBytes []byte, kid string) error {
	var err error

	if len(privateKeyBytes) > 0 {
		c.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
		if err != nil {
			return err
		}
	}

	if len(publicKeyBytes) > 0 {
		c.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
		if err != nil {
			return err
		}
	}

	c.kid = kid

	err = c.Jwk.CreateJwks()
	if err != nil {
		return err
	}

	return nil
}

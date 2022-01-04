package core

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/mechta-market/jwts/internal/domain/entities"
)

type Jwk struct {
	r *St

	jwks *entities.JwkSetSt
}

func NewJwk(r *St) *Jwk {
	return &Jwk{
		r: r,
	}
}

func (c *Jwk) CreateJwks() error {
	var err error

	c.jwks, err = c.createJwks()
	if err != nil {
		return err
	}

	return nil
}

func (c *Jwk) createJwks() (*entities.JwkSetSt, error) {
	if c.r.publicKey == nil {
		return nil, nil
	}

	eBytes := make([]byte, 4, 4)
	binary.LittleEndian.PutUint32(eBytes, uint32(c.r.publicKey.E))

	key := entities.JwkSt{
		Kty: "RSA",
		E:   base64.RawURLEncoding.EncodeToString(eBytes[:3]),
		Kid: c.r.kid,
		Alg: "RS256",
		N:   base64.RawURLEncoding.EncodeToString(c.r.publicKey.N.Bytes()),
		Use: "sig",
	}

	return &entities.JwkSetSt{
		Keys: []*entities.JwkSt{
			&key,
		},
	}, nil
}

func (c *Jwk) GetSet() *entities.JwkSetSt {
	return c.jwks
}

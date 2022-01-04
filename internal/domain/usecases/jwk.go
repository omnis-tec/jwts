package usecases

import (
	"github.com/mechta-market/jwts/internal/domain/entities"
)

func (u *St) JwkGetSet() *entities.JwkSetSt {
	return u.cr.Jwk.GetSet()
}

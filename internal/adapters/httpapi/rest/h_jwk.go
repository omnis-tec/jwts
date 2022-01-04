package rest

import (
	"net/http"

	"github.com/mechta-market/jwts/internal/domain/entities"
)

// swagger:route GET /jwk/set jwk hJwkGetSet
// Список публичных ключей jwk.
// Responses:
//   200: jwkGetSetRep
//   400: errRep
func (a *St) hJwkGetSet(w http.ResponseWriter, r *http.Request) {
	// swagger:response jwkGetSetRep
	type docRepSt struct {
		// in:body
		Body *entities.JwkSetSt
	}

	repObj := a.ucs.JwkGetSet()

	a.uRespondJSON(w, repObj)
}

package rest

import (
	"net/http"

	"github.com/mechta-market/jwts/internal/domain/entities"
)

// swagger:route POST /jwt jwt hJwtCreate
// Создать jwt-токен.
// Responses:
//   200: jwtCreateRep
//   400: errRep
func (a *St) hJwtCreate(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters hJwtCreate
	type docReqSt struct {
		// in: body
		Body entities.JwtCreateReqSt
	}

	// swagger:response jwtCreateRep
	type docRepSt struct {
		// in:body
		Body struct {
			Token string `json:"token"`
		}
	}

	reqObj := map[string]interface{}{}
	if !a.uParseRequestJSON(w, r, &reqObj) {
		return
	}

	token, err := a.ucs.JwtCreate(reqObj)
	if a.uHandleError(err, r, w) {
		return
	}

	a.uRespondJSON(w, struct {
		Token string `json:"token"`
	}{
		Token: token,
	})
}

// swagger:route PUT /jwt/validate jwt hJwtValidate
// Проверить и распарсить jwt-токен.
// Responses:
//   200: jwtValidateRep
//   400: errRep
func (a *St) hJwtValidate(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters hJwtValidate
	type docReqSt struct {
		// in: body
		Body struct {
			Token string `json:"token"`
		}
	}

	// swagger:response jwtValidateRep
	type docRepSt struct {
		// in:body
		Body *entities.JwtValidateRepSt
	}

	reqObj := &(struct {
		Token string `json:"token"`
	}{})
	if !a.uParseRequestJSON(w, r, reqObj) {
		return
	}

	repObj, err := a.ucs.JwtValidate(reqObj.Token)
	if a.uHandleError(err, r, w) {
		return
	}

	a.uRespondJSON(w, repObj)
}

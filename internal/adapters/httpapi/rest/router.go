/*
Package rest Jwts-service API.


    Schemes: https, http
    Host: jwts.com
    BasePath: /
    Version: 1.0.0

    Consumes:
    - application/json

    Produces:
    - application/json

    SecurityDefinitions:
      token:
         type: apiKey
         name: Authorization
         in: header
         description: "Пример: `Authorization: 2cf24dba5fb0a30e26e83b2`"

swagger:meta
*/
package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *St) router() http.Handler {
	r := mux.NewRouter()

	// doc
	r.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "doc/")
		w.WriteHeader(http.StatusMovedPermanently)
	})
	r.PathPrefix("/doc/").Handler(http.StripPrefix("/doc/", http.FileServer(http.Dir("./doc/"))))

	// jwk
	r.HandleFunc("/jwk/set", a.hJwkGetSet).Methods("GET")

	// jwt
	r.HandleFunc("/jwt", a.hJwtCreate).Methods("POST")
	r.HandleFunc("/jwt/validate", a.hJwtValidate).Methods("PUT")

	return a.middleware(r)
}

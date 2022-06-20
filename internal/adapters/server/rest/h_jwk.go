package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router   /jwk/set [get]
// @Tags     jwk
// @Produce  json
// @Success  200  {object}  entities.JwkSetSt
// @Failure  400  {object}  dopTypes.ErrRep
func (a *St) hJwkGetSet(c *gin.Context) {
	c.JSON(http.StatusOK, a.core.Jwk.GetSet())
}

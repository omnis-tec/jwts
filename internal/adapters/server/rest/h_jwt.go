package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/jwts/internal/domain/entities"
)

// @Router   /jwt [post]
// @Tags     jwt
// @Param    body  body  entities.JwtCreateReqSt  false  "body"
// @Produce  json
// @Success  200  {object}  entities.JwtCreateRepSt
// @Failure  400  {object}  dopTypes.ErrRep
func (a *St) hJwtCreate(c *gin.Context) {
	reqObj := map[string]any{}
	if !dopHttps.BindJSON(c, &reqObj) {
		return
	}

	result, err := a.core.Jwt.Create(reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /jwt/validate [put]
// @Tags     jwt
// @Param    body  body  entities.JwtValidateReqSt  false  "body"
// @Produce  json
// @Success  200  {object}  entities.JwtValidateRepSt
// @Failure  400  {object}  dopTypes.ErrRep
func (a *St) hJwtValidate(c *gin.Context) {
	reqObj := &entities.JwtValidateReqSt{}
	if !dopHttps.BindJSON(c, &reqObj) {
		return
	}

	result, err := a.core.Jwt.Validate(reqObj.Token)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

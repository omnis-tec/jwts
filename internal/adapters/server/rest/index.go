package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/logger"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/jwts/docs"
	"github.com/rendau/jwts/internal/domain/core"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

type St struct {
	lg   logger.Lite
	core *core.St
}

func GetHandler(lg logger.Lite, core *core.St, withCors bool) http.Handler {
	docs.SwaggerInfo.Host = "jwts.jwts"
	docs.SwaggerInfo.BasePath = "/jwts"
	docs.SwaggerInfo.Schemes = []string{"https"}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(dopHttps.MwRecovery(lg, nil))
	if withCors {
		r.Use(dopHttps.MwCors())
	}

	// handlers

	// doc
	r.GET("/doc/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))

	s := &St{lg: lg, core: core}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// jwk
	r.GET("/jwk/set", s.hJwkGetSet)

	// jwt
	r.POST("/jwt", s.hJwtCreate)
	r.PUT("/jwt/validate", s.hJwtValidate)

	return r
}

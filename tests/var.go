package tests

import (
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	"github.com/rendau/jwts/internal/domain/core"
)

var (
	app = struct {
		lg   *dopLoggerZap.St
		core *core.St
	}{}
)

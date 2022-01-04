package tests

import (
	"github.com/mechta-market/jwts/internal/adapters/logger/zap"
	"github.com/mechta-market/jwts/internal/domain/core"
	"github.com/mechta-market/jwts/internal/domain/usecases"
)

var (
	app = struct {
		lg   *zap.St
		core *core.St
		ucs  *usecases.St
	}{}
)

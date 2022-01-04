package usecases

import (
	"github.com/mechta-market/jwts/internal/domain/core"
	"github.com/mechta-market/jwts/internal/interfaces"
)

type St struct {
	lg interfaces.Logger

	cr *core.St
}

func New(
	lg interfaces.Logger,
	cr *core.St,
) *St {
	u := &St{
		lg: lg,
		cr: cr,
	}

	return u
}

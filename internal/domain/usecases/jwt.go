package usecases

import "github.com/mechta-market/jwts/internal/domain/entities"

func (u *St) JwtCreate(reqObj map[string]interface{}) (string, error) {
	return u.cr.Jwt.Create(reqObj)
}

func (u *St) JwtValidate(value string) (*entities.JwtValidateRepSt, error) {
	return u.cr.Jwt.Validate(value)
}

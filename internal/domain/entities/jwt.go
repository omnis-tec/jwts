package entities

type JwtCreateReqSt struct {
	Sub        string         `json:"sub"`
	ExpSeconds int64          `json:"exp_seconds"`
	Payload    map[string]any `json:"payload"`
}

type JwtCreateRepSt struct {
	Token string `json:"token"`
}

type JwtValidateReqSt struct {
	Token string `json:"token"`
}

type JwtValidateRepSt struct {
	Valid  bool           `json:"valid"`
	Claims map[string]any `json:"claims"`
}

package entities

type JwtCreateReqSt struct {
	Sub        string                 `json:"sub"`
	ExpSeconds int64                  `json:"exp_seconds"`
	Payload    map[string]interface{} `json:"payload"`
}

type JwtValidateRepSt struct {
	Valid  bool                   `json:"valid"`
	Claims map[string]interface{} `json:"claims"`
}

package entities

type JwkSt struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	Use string `json:"use"`
}

type JwkSetSt struct {
	Keys []*JwkSt `json:"keys"`
}

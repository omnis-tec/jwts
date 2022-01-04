package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	BadJson      = Err("bad_json")
	ServiceNA    = Err("server_not_available")
	InvalidToken = Err("invalid_token")
)

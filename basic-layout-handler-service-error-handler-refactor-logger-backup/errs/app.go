package errs

type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

func NotFound(msg string) *Error {
	return &Error{
		Code:    404,
		Message: msg,
	}
}

func ServerError(msg string) *Error {
	return &Error{
		Code:    502,
		Message: msg,
	}
}

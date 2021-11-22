package errs

import "net/http"

// AppErr
type AppErr struct {
	Code    int
	Message string
}

// NotFound ...
func NotFound() *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
}

// ServerError ...
func ServerError() *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

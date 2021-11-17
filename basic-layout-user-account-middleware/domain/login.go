package domain

import (
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
)

type TokenResponse struct {
	TokenString string `json:"token"`
}

type Login struct {
	Email    string
	Password string
}

type LoginRepository interface {
	FindByEmail(LoginRequest) (*Login, *errs.Error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

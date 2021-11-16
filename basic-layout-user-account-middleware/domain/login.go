package domain

import (
	"time"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
)

const TOKEN_EXPIRE = time.Hour
const HMAC_SAMPLE_SECRET = "hmacSampleSECRET"

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

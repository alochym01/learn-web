package domain

import (
	"time"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
)

const TOKEN_EXPIRE = time.Hour
const HMAC_SAMPLE_SECRET = "hmacSampleSECRET"

type Claim struct {
	Expire int
	Token  string
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

func (l Login) GenerateToken() (*string, error) {

}

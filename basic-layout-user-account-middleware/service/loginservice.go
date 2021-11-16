package service

import (
	"time"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
	"github.com/golang-jwt/jwt"
)

// LoginService ...
type LoginService struct {
	storageRepo domain.LoginRepository
}

// LoginServiceRepository ...
type LoginServiceRepository interface {
	ByEmail(domain.LoginRequest) (*domain.TokenResponse, *errs.Error)
}

// NewLoginService ...
func NewLoginService(repo domain.LoginRepository) LoginService {
	return LoginService{
		storageRepo: repo,
	}
}

// ByEmail ...
func (ls LoginService) ByEmail(temp domain.LoginRequest) (*domain.TokenResponse, *errs.Error) {
	// 1. Check user/pass are valid
	u, err := ls.storageRepo.FindByEmail(temp)
	if err != nil {
		return nil, err
	}

	// 2. Generating a token
	// Create jwt MapClaims
	claims := jwt.MapClaims{}
	claims["user"] = u.Email
	claims["exp"] = time.Now().Add(domain.TOKEN_EXPIRE).Unix()

	// Create jwt Claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err1 := token.SignedString([]byte(domain.HMAC_SAMPLE_SECRET))
	if err1 != nil {
		return nil, errs.ServerError("Server Internal Error")
	}

	t := domain.TokenResponse{}
	t.TokenString = tokenString
	// 3. Return a Token String
	return &t, nil
}

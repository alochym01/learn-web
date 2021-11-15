package service

import (
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
)

// LoginService ...
type LoginService struct {
	storageRepo domain.LoginRepository
}

// LoginServiceRepository ...
type LoginServiceRepository interface {
	ByEmail(domain.LoginRequest) (*domain.Login, *errs.Error)
}

// NewLoginService ...
func NewLoginService(repo domain.LoginRepository) LoginService {
	return LoginService{
		storageRepo: repo,
	}
}

// ByEmail ...
func (ls LoginService) ByEmail(temp domain.LoginRequest) (*domain.Login, *errs.Error) {
	return ls.storageRepo.FindByEmail(temp)
}

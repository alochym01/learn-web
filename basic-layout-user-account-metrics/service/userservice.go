package service

import (
	"fmt"
	"strconv"

	"github.com/alochym01/learn-web/basic-layout-user-account/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account/errs"
)

// UserService ...
type UserService struct {
	storageRepo domain.UserRepository
}

// UserServiceRepository ...
type UserServiceRepository interface {
	GetUsers() ([]domain.User, *errs.Error)
	// ByEmail(string) (*domain.User, *errs.Error)
	ByID(string) (*domain.User, *errs.Error)
	Create(domain.UserRequest) *errs.Error
	Update(string, domain.UserRequest) *errs.Error
	Delete(string) *errs.Error
}

// NewUserService ...
func NewUserService(repo domain.UserRepository) UserService {
	return UserService{
		storageRepo: repo,
	}
}

// GetUsers ...
func (us UserService) GetUsers() ([]domain.User, *errs.Error) {
	return us.storageRepo.FindAll()
}

// ByID ...
func (us UserService) ByID(temp string) (*domain.User, *errs.Error) {
	return us.storageRepo.FindByID(temp)
}

// ByEmail ...
func (us UserService) ByEmail(temp string) (*domain.User, *errs.Error) {
	return us.storageRepo.FindByEmail(temp)
}

// Create ...
func (us UserService) Create(temp domain.UserRequest) *errs.Error {
	var u domain.User
	u.Email = temp.Email
	u.Password = temp.Password
	u.FullName = temp.FullName
	fmt.Println(u)
	return us.storageRepo.Create(u)
}

// Update ...
func (us UserService) Update(id string, temp domain.UserRequest) *errs.Error {
	var u domain.User
	u.Email = temp.Email
	u.Password = temp.Password
	u.FullName = temp.FullName
	u.ID, _ = strconv.Atoi(id)
	return us.storageRepo.Update(u)
}

// Delete ...
func (us UserService) Delete(id string) *errs.Error {
	return us.storageRepo.Delete(id)
}

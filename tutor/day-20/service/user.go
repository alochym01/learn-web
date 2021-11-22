package service

import (
	"fmt"
	"strconv"

	"github.com/alochym01/web-w-gin/domain"
	"github.com/alochym01/web-w-gin/errs"
)

// UserService ...
type UserService struct {
	storageRepo domain.UserRepository
}

// UserServiceRepository ...
type UserServiceRepository interface {
	GetUsers() ([]domain.User, *errs.AppErr)
	ByEmail(string) (*domain.User, *errs.AppErr)
	ByID(int) (*domain.User, *errs.AppErr)
	Create(domain.UserRequest) (*int64, *errs.AppErr)
}

// NewUserService ...
func NewUserService(repo domain.UserRepository) UserService {
	return UserService{
		storageRepo: repo,
	}
}

// GetUsers ...
func (us UserService) GetUsers() ([]domain.User, *errs.AppErr) {
	return us.storageRepo.FindAll()
}

// ByID ...
func (us UserService) ByID(temp int) (*domain.User, *errs.AppErr) {
	return us.storageRepo.FindByID(temp)
}

// ByEmail ...
func (us UserService) ByEmail(temp string) (*domain.User, *errs.AppErr) {
	return us.storageRepo.FindByEmail(temp)
}

// Create ...
func (us UserService) Create(temp domain.UserRequest) (*int64, *errs.AppErr) {
	var u domain.User
	u.Email = temp.Email
	u.Password = temp.Password
	u.FullName = temp.FullName
	fmt.Println(u)
	return us.storageRepo.Create(u)
}

// Update ...
func (us UserService) Update(id string, temp domain.UserRequest) *errs.AppErr {
	var u domain.User
	u.Email = temp.Email
	u.Password = temp.Password
	u.FullName = temp.FullName
	u.ID, _ = strconv.Atoi(id)
	return us.storageRepo.Update(u)
}

// Delete ...
func (us UserService) Delete(id string) *errs.AppErr {
	return us.storageRepo.Delete(id)
}

package domain

import (
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
)

type User struct {
	ID       int
	FullName string
	Email    string
	Password string
}

type UserRepository interface {
	FindAll() ([]User, *errs.Error)
	FindByEmail(string) (*User, *errs.Error)
	FindByID(string) (*User, *errs.Error)
	Create(User) *errs.Error
	Update(User) *errs.Error
	Delete(string) *errs.Error
}

type UserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

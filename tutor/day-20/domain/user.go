package domain

import (
	"github.com/alochym01/web-w-gin/errs"
)

type User struct {
	ID       int
	FullName string
	Email    string
	Password string
}

type UserRepository interface {
	FindAll() ([]User, *errs.AppErr)
	FindByEmail(string) (*User, *errs.AppErr)
	FindByID(int) (*User, *errs.AppErr)
	Create(User) (*int64, *errs.AppErr)
	Update(User) *errs.AppErr
	Delete(string) *errs.AppErr
}

type UserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

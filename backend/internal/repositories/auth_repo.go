package repository

import (
	"s-store/internal/model/request"
	"s-store/internal/model/response"

	"gorm.io/gorm"
)

type AuthRepo interface {
	ResgisterUser(user request.RegisterRequest) (interface{}, error)
	LoginUser(loginReq request.LoginRequest) (response.AuthResponse, error)
}

type authRepo struct {
	db *gorm.DB
}

// LoginUser implements [AuthRepo].
func (a *authRepo) LoginUser(loginReq request.LoginRequest) (response.AuthResponse, error) {
	panic("unimplemented")
}

// ResgisterUser implements [AuthRepo].
func (a *authRepo) ResgisterUser(user request.RegisterRequest) (interface{}, error) {
	panic("unimplemented")
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{
		db: db,
	}
}

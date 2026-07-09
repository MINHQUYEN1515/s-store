package service

import (
	"s-store/internal/model/request"
	"s-store/internal/model/response"
	repository "s-store/internal/repositories"
)

type AuthService interface {
	Login(request request.LoginRequest) (response.AuthResponse, error)
	Register(request request.RegisterRequest) (interface{}, error)
}

type authService struct {
	authRepo  repository.AuthRepo
	jwtSecret string
}

// Login implements [AuthService].
func (a *authService) Login(request request.LoginRequest) (response.AuthResponse, error) {
	panic("unimplemented")
}

// Register implements [AuthService].
func (a *authService) Register(request request.RegisterRequest) (interface{}, error) {
	panic("unimplemented")
}

func AuthServiceNew(authRepo repository.AuthRepo, jwtSecret string) AuthService {
	return &authService{
		authRepo:  authRepo,
		jwtSecret: jwtSecret,
	}
}

package service

import (
	"errors"
	"s-store/internal/model/entity"
	"s-store/internal/model/request"
	"s-store/internal/model/response"
	repository "s-store/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(request request.LoginRequest) (response.AuthResponse, error)
	Register(request request.RegisterRequest) (*entity.UserEntity, error)
}

type authService struct {
	userRepo  repository.UserRepo
	jwtSecret string
}

// Login implements [AuthService].
func (a *authService) Login(request request.LoginRequest) (*response.AuthResponse, error) {
	panic("unimplemented")
}

// Register implements [AuthService].
func (a *authService) Register(request request.RegisterRequest) (*entity.UserEntity, error) {
	existingUser, err := a.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("Email already exists")
	}
	hashPassword, err := hashPassword(request.Password)
	user := &entity.UserEntity{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	}
	user, err = a.userRepo.CreaterUser(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthServiceNew(UserRepo repository.UserRepo, jwtSecret string) AuthService {
	return &authService{
		userRepo:  UserRepo,
		jwtSecret: jwtSecret,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

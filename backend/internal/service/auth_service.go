package service

import (
	"errors"
	"s-store/internal/model/entity"
	"s-store/internal/model/enum"
	"s-store/internal/model/request"
	"s-store/internal/model/response"
	repository "s-store/internal/repositories"
	jwtpackage "s-store/pkg/jwt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(request request.LoginRequest) (*response.AuthResponse, error)
	Register(request request.RegisterRequest) (*entity.UserEntity, error)
}

type authService struct {
	userRepo        repository.UserRepo
	jwtSecret       string
	jwtSecretRefesh string
}

// Login implements [AuthService].
func (a *authService) Login(request request.LoginRequest) (*response.AuthResponse, error) {
	user, err := a.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}
	token, err := jwtpackage.GenerateToken(jwtpackage.AuthKey{
		UserId:    user.ID,
		Role:      user.Role,
		SecretKey: a.jwtSecret,
		TimeExp:   time.Now().Add(12 * time.Hour), // Token expiration time
		Type:      enum.TokenTypeAccess,
	})

	if err != nil {
		return nil, err
	}

	refeshToken, err := jwtpackage.GenerateToken(jwtpackage.AuthKey{
		UserId:    user.ID,
		Role:      user.Role,
		SecretKey: a.jwtSecretRefesh,
		TimeExp:   time.Now().Add(30 * 24 * time.Hour), // Refesh Token expiration time
		Type:      enum.TokenTypeRefresh,
	})

	if err != nil {
		return nil, err
	}

	return &response.AuthResponse{
		Token:       token,
		RefeshToken: refeshToken,
		User:        *user,
	}, nil
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

func AuthServiceNew(UserRepo repository.UserRepo, jwtSecret string, jwtSecretRefesh string) AuthService {
	return &authService{
		userRepo:        UserRepo,
		jwtSecret:       jwtSecret,
		jwtSecretRefesh: jwtSecretRefesh,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

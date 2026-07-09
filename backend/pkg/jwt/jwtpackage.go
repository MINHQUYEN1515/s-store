package jwtpackage

import (
	"errors"
	backend "s-store"
	"s-store/internal/model/enum"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role enum.RoleEnum, config *backend.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour).Unix(), // Token expires in 1 hours,
		"iat":     time.Now().Unix(),                // Token issued at
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWTKey))
}

func ValidateToken(tokenString string, config *backend.Config) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.JWTKey), nil
	})
	if err != nil {
		return nil, err
	}
	// Check token hợp lệ + lấy claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

package jwtpackage

import (
	backend "s-store"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, config *backend.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"EZ"
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWTKey))
}

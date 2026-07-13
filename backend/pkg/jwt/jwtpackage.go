package jwtpackage

import (
	"errors"
	"s-store/internal/model/enum"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthKey struct {
	UserId    uint
	Role      enum.RoleEnum
	SecretKey string
	TimeExp   time.Time
	Type      enum.TokenType
}

func GenerateToken(authKey AuthKey) (string, error) {
	claims := jwt.MapClaims{
		"user_id": authKey.UserId,
		"role":    authKey.Role,
		"exp":     authKey.TimeExp.Unix(),
		"iat":     time.Now().Unix(), // Token issued at
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(authKey.SecretKey))
}

func ValidateToken(tokenString string, jwtSecret string, tokenType enum.TokenType) (*AuthKey, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	// Check token hợp lệ + lấy claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims["type"] != string(tokenType) {
		return nil, errors.New("invalid token type")
	}
	return &AuthKey{
		UserId:    uint(claims["user_id"].(float64)),
		Role:      enum.RoleEnum(claims["role"].(string)),
		SecretKey: jwtSecret,
		TimeExp:   time.Unix(int64(claims["exp"].(float64)), 0),
		Type:      tokenType,
	}, nil
}

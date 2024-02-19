package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	CreateAccessToken(id, fullname string) (string, error)
	VerifyToken(tokenStr string) (*jwt.MapClaims, error)
	GetIdAndExp(claims *jwt.MapClaims) (string, int64, error)
}

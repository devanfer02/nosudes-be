package service

import (
	"time"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/domain"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
}

func NewAuthService() domain.AuthService {
	return &authService{}
}

func (s *authService) CreateAccessToken(id, fullname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"name": fullname,
		"iss":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(env.ProcEnv.ATSecret))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *authService) VerifyToken(tokenStr string) (*jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.ProcEnv.ATSecret), nil
	})

	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	if !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, domain.ErrInvalidToken
	}

	return &claims, nil
}

func (s *authService) GetIdAndExp(claims *jwt.MapClaims) (string, int64, error) {
	exp, ok := (*claims)["exp"].(float64)

	if !ok {
		
		return "", 0, domain.ErrInvalidClaimsDT
	}

	id, ok := (*claims)["id"].(string)

	if !ok {

		return "", 0, domain.ErrInvalidClaimsDT
	}

	return id, int64(exp), nil
}

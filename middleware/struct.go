package middleware

import "github.com/devanfer02/nosudes-be/domain"

type Middleware struct {
	userSvc domain.UserService
	authSvc domain.AuthService
}

func NewMiddleware(userSvc domain.UserService, authSvc domain.AuthService) *Middleware {
	return &Middleware{userSvc, authSvc}
}
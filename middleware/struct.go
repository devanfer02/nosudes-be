package middleware

import "github.com/devanfer02/nosudes-be/domain"

type Middleware struct {
	userSvc domain.UserService
}

func NewMiddleware(userSvc domain.UserService) *Middleware {
	return &Middleware{userSvc}
}
package service

import (
	"context"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

type userService struct {
	repo domain.UserRepository
	timeout time.Duration
}

func NewUserService(repo domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{repo, timeout}
}

func (s *userService) FetchAll(ctx context.Context) ([]domain.User, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)	
	defer cancel()

	users, err := s.repo.FetchAll(c)

	return users, err
}

func (s *userService) FetchByEmail(ctx context.Context, email string) (domain.User, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.repo.FetchOneByArg(c, "email", email)

	return user, err 
}

func (s *userService) FetchByID(ctx context.Context, id string) (domain.User, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.repo.FetchOneByArg(c, "id", id)

	return user, err 
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.repo.UpdateUser(c, user)

	return err 
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.repo.DeleteUser(c, id)

	return err 
}

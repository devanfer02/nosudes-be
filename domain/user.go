package domain

import (
	"context"
)

type User struct {
	ID        string `json:"id" db:"user_id"`
	Fullname  string `json:"fullname" db:"fullname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type UserRepository interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchOneByArg(ctx context.Context, param, arg string) (User, error)
	InsertUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
}

type UserService interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchByEmail(ctx context.Context, email string) (User, error)
	FetchByID(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
}

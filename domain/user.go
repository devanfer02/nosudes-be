package domain

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id" db:"user_id"`
	Fullname  string `json:"fullname" db:"fullname"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type UserPayload struct {
	ID       string
	Fullname string `json:"id" valid:"required"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required, stringlength(8|50)"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required"`
}

type UserRepository interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchOneByArg(ctx context.Context, param, arg string) (User, error)
	InsertUser(ctx context.Context, user *UserPayload) error
	UpdateUser(ctx context.Context, user *UserPayload) error
	DeleteUser(ctx context.Context, id string) error
}

type UserService interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchByEmail(ctx context.Context, email string) (User, error)
	FetchByID(ctx context.Context, id string) (User, error)
	InsertUser(ctx context.Context, user *UserPayload) error
	UpdateUser(ctx context.Context, user *UserPayload) error
	DeleteUser(ctx context.Context, id string) error
}

func (p *UserPayload) Default() {
	byteHashed, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

	hashed := string(byteHashed)

	p.Password = hashed
	p.ID = uuid.New().String()
}

func (u *User) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); 

	return err != nil 
}

package domain

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       		string 	`json:"id" db:"user_id"`
	Fullname 		string 	`json:"fullname" db:"fullname"`
	Username 		string 	`json:"username" db:"username"`
	Email    		string 	`json:"email" db:"email"`
	Password 		string 	`json:"-" db:"password"`
	PPURL    		string 	`json:"photo_profile" db:"photo_url"`
	TotalLikes		int64	`json:"total_likes" db:"total_likes"`
	TotalReviews	int64	`json:"total_reviews" db:"total_reviews"`
	CreatedAt 		string 	`json:"created_at" db:"created_at"`
	UpdatedAt 		string 	`json:"updated_at" db:"updated_at"`
}

type UserPayload struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname" valid:"required"`
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"required, stringlength(8|50)"`
}

type UserPhotoPayload struct {
	UserID       string                `form:"-"`
	PhotoProfile *multipart.FileHeader `form:"photo_profile" binding:"required"`
	PhotoURL     string                `form:"-"`
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
	UpdatePP(ctx context.Context, photo *UserPhotoPayload) error
	DeleteUser(ctx context.Context, id string) error
}

type UserService interface {
	FetchAll(ctx context.Context) ([]User, error)
	FetchByEmail(ctx context.Context, email string) (User, error)
	FetchByID(ctx context.Context, id string) (User, error)
	InsertUser(ctx context.Context, user *UserPayload) error
	UploadPP(ctx context.Context, photo *UserPhotoPayload) error
	UpdateUser(ctx context.Context, user *UserPayload) error
	DeleteUser(ctx context.Context, id string) error
}

func (p *UserPayload) Default() {
	byteHashed, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

	hashed := string(byteHashed)

	p.Password = hashed
	p.ID = uuid.New().String()
}

func (p *UserPayload) DefaultWithID() {
	byteHashed, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)

	hashed := string(byteHashed)

	p.Password = hashed
}

func (u *User) Compare(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

package domain

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          string    				`json:"id" db:"article_id"`
	Title       string    				`json:"title" db:"title" binding:"required"`
	Summary     string    				`json:"summary" db:"summary" binding:"required"`
	Description string    				`json:"description" db:"description" binding:"required"`
	Photo       string    				`json:"photo" db:"photo"`
	CreatedAt   time.Time 				`json:"created_at" db:"created_at"`
}

type ArticlePayload struct {
	ID			string
	Title       string    				`form:"title" binding:"required"`
	Summary     string    				`form:"summary" binding:"required"`
	Description string    				`form:"description" binding:"required"`
	PhotoFile   *multipart.FileHeader   `form:"photo" binding:"required"`
}

type ArticleRepository interface {
	FetchAll(ctx context.Context) ([]Article, error)
	FetchByID(ctx context.Context, id string) (Article, error)
	InsertArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, article *Article) error 
	DeleteArticle(ctx context.Context, id string) error
}

type ArticleService interface {
	FetchAll(ctx context.Context) ([]Article, error)
	FetchByID(ctx context.Context, id string) (Article, error)
	InsertArticle(ctx context.Context, article *ArticlePayload) error
	UpdateArticle(ctx context.Context, article *ArticlePayload) error 
	DeleteArticle(ctx context.Context, id string) error
}

func(a *ArticlePayload) Convert(photoUrl string) *Article {
	return &Article{
		ID: uuid.New().String(),
		Title: a.Title,
		Summary: a.Summary,
		Description: a.Description,
		Photo: photoUrl,
	}
}
package domain

import (
	"context"
	"time"
	"mime/multipart"

	"github.com/google/uuid"
)

type Review struct {
	ID           string `json:"review_id" db:"review_id"`
	AttractionID string `json:"-" db:"attraction_id"`
	UserID       string `json:"-" db:"user_id"`
	ReviewText   string `json:"review_text" db:"review_text"`
	PhotoURL     string `json:"photo_url" db:"photp_url"`
	DateCreated  string `json:"date_created" db:"date_created"`
}

type ReviewPayload struct {
	ID           string 				`db:"review_id"`
	AttractionID string 				`json:"-" db:"attraction_id"`
	UserID       string 				`json:"-" db:"user_id"`
	ReviewText   string 				`form:"review_text" binding:"required" valid:"required, stringlength(2|400)" db:"review_text"`
	PhotoURL     string 				`form:"-" db:"photp_url"`
	DateCreated  string 				`form:"-" db:"date_created"`
	PhotoFile	 *multipart.FileHeader	`form:"photo" db:"-"`
}

type ReviewRepository interface {
	FetchAll(ctx context.Context) ([]*Review, error)
	FetchByAttrID(ctx context.Context, attractionId string) ([]*Review, error)
	FetchByID(ctx context.Context, id string) (*Review, error)
	InsertReview(ctx context.Context, review *ReviewPayload) error
	DeleteReview(ctx context.Context, reviewId string) error
}

type ReviewService interface {
	FetchAll(ctx context.Context) ([]*Review, error)
	FetchByAttrID(ctx context.Context, attractionId string) ([]*Review, error)
	FetchByID(ctx context.Context, id string) (*Review, error)
	InsertReview(ctx context.Context, review *ReviewPayload) error
	DeleteReview(ctx context.Context, reviewId, userId string) error
}

func (rp *ReviewPayload) Default(attractionId, userId string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	rp.DateCreated = time.Now().In(loc).Format("2006-01-02")
	rp.ID = uuid.New().String()
	rp.AttractionID = attractionId
	rp.UserID = userId
}

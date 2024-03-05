package domain

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID           string     `json:"review_id" db:"review_id"`
	AttractionID string     `json:"attraction_id" db:"attraction_id"`
	UserID       string     `json:"user_id" db:"user_id"`
	ReviewText   string     `json:"review_text" db:"review_text"`
	Liked		 int		`json:"liked" db:"liked"`
	PhotoURL     string     `json:"photo_url" db:"photo_url"`
	DateCreated  string     `json:"date_created" db:"date_created"`
	UserDetail   User       `json:"user_detail" db:"-"`
	AttrDetail   Attraction `json:"attraction_detail" db:"-"`
}

type ReviewPayload struct {
	ID           string                `db:"review_id" json:"id"`
	AttractionID string                `json:"attraction_id" db:"attraction_id" form:"-"`
	UserID       string                `json:"user_id" db:"user_id" form:"-"`
	ReviewText   string                `form:"review_text" binding:"required" valid:"required, stringlength(2|400)" db:"review_text" json:"review_text"`
	PhotoURL     string                `form:"-" db:"photo_url" json:"photo_url"`
	DateCreated  string                `form:"-" db:"date_created" json:"date_created"`
	PhotoFile    *multipart.FileHeader `form:"photo" db:"-" json:"-"`
}

type ReviewRepository interface {
	FetchAll(ctx context.Context, args ...interface{}) ([]*Review, error)
	FetchByAttrID(ctx context.Context, attractionId string, args ...interface{}) ([]*Review, error)
	FetchByID(ctx context.Context, id string, args ...interface{}) (*Review, error)
	LikeReview(ctx context.Context, reviewId, userId string) error
	UnlikeReview(ctx context.Context, reviewId, userId string) error 
	InsertReview(ctx context.Context, review *ReviewPayload) error
	DeleteReview(ctx context.Context, reviewId string) error
}

type ReviewService interface {
	FetchAll(ctx context.Context, args ...interface{}) ([]*Review, error)
	FetchByAttrID(ctx context.Context, attractionId string, args ...interface{}) ([]*Review, error)
	FetchByID(ctx context.Context, id string, args ...interface{}) (*Review, error)
	LikeReview(ctx context.Context, reviewId, userId string) error
	UnlikeReview(ctx context.Context, reviewId, userId string) error 
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

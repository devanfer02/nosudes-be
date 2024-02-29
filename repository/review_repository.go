package repository

import (
	"context"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"

	"github.com/jmoiron/sqlx"
)

type mysqlReviewRepository struct {
	Conn *sqlx.DB
}

func NewMysqlReviewRepository(conn *sqlx.DB) domain.ReviewRepository {
	return &mysqlReviewRepository{conn}
}

func (m *mysqlReviewRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*domain.Review, error) {
	reviews := make([]*domain.Review, 0)

	err := m.Conn.SelectContext(ctx, &reviews, query, args...)

	if err != nil {
		logger.ErrLog(layers.Repository, "failed to query", err)
		return nil, err
	}

	return reviews, err
}

func (m *mysqlReviewRepository) FetchAll(ctx context.Context) ([]*domain.Review, error) {
	query := `SELECT * FROM reviews`

	reviews, err := m.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (m *mysqlReviewRepository) FetchByAttrID(ctx context.Context, attractionId string) ([]*domain.Review, error) {
	query := `SELECT * FROM reviews WHERE attraction_id = ?`

	reviews, err := m.fetch(ctx, query, attractionId)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (m *mysqlReviewRepository) FetchByID(ctx context.Context, id string) (*domain.Review, error) {
	query := `SELECT * FROM reviews WHERE review_id = ?`

	reviews, err := m.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, domain.ErrNotFound
	}

	return reviews[0], nil
}

func (m *mysqlReviewRepository) LikeReview(ctx context.Context, reviewId, userId string) error {
	query := `INSERT INTO review_likes (review_id, user_id) VALUES (?, ?)`

	return execStatement(m.Conn, ctx, query, reviewId, userId)
}

func (m *mysqlReviewRepository) UnlikeReview(ctx context.Context, reviewId, userId string) error {
	query := `DELETE FROM review_likes WHERE review_id = ? AND user_id = ?`

	return execStatement(m.Conn, ctx, query, reviewId, userId)
}

func (m *mysqlReviewRepository) InsertReview(ctx context.Context, review *domain.ReviewPayload) error {

	query := `INSERT INTO reviews 
	(review_id, attraction_id, user_id, review_text, photo_url, date_created) 
	VALUES (?, ?, ?, ?, ?, ?)`

	return execStatement(
		m.Conn,
		ctx,
		query,
		review.ID, review.AttractionID,
		review.UserID, review.ReviewText, review.PhotoURL,
		review.DateCreated,
	)
}

func (m *mysqlReviewRepository) DeleteReview(ctx context.Context, reviewId string) error {
	query := `DELETE FROM reviews WHERE review_id = ?`

	return execStatement(
		m.Conn,
		ctx,
		query,
		reviewId,
	)
}

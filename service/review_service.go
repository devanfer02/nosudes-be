package service

import (
	"context"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

const REVIEWS_CLOUD_STORE_DIR = "reviews"

type reviewService struct {
	rvRepo domain.ReviewRepository
	fileStore domain.FileStorage
	timeout time.Duration
}

func NewReviewService(rvRepo domain.ReviewRepository, fileStore domain.FileStorage, timeout time.Duration) domain.ReviewService {
	return &reviewService{rvRepo, fileStore, timeout}
}

func(s *reviewService) FetchAll(ctx context.Context) ([]*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	reviews, err := s.rvRepo.FetchAll(c) 

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func(s *reviewService) FetchByAttrID(ctx context.Context, attractionId string) ([]*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	reviews, err := s.rvRepo.FetchByAttrID(c, attractionId) 

	if err != nil {
		return nil, err
	}

	return reviews, nil

}
func(s *reviewService) FetchByID(ctx context.Context, id string) (*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	review, err := s.rvRepo.FetchByID(c, id) 

	if err != nil {
		return nil, err
	}

	return review, nil
}

func(s *reviewService) InsertReview(ctx context.Context, review *domain.ReviewPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if review.PhotoFile != nil {
		url, err := s.fileStore.UploadFile(REVIEWS_CLOUD_STORE_DIR, review.PhotoFile)

		if err != nil {
			return err
		}

		review.PhotoURL = url
	}

	err := s.rvRepo.InsertReview(c, review)

	return err
}

func(s *reviewService) DeleteReview(ctx context.Context, reviewId, userId string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	review, err := s.rvRepo.FetchByID(c, reviewId)

	if err != nil {
		return err 
	}

	if review.UserID != userId {
		return domain.ErrForbidden
	}

	err = s.rvRepo.DeleteReview(c, reviewId)

	return err
}

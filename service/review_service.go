package service

import (
	"context"
	"sync"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

const REVIEWS_CLOUD_STORE_DIR = "reviews"

type reviewService struct {
	rvRepo    domain.ReviewRepository
	usrRepo   domain.UserRepository
	attrRepo  domain.AttractionRepository
	fileStore domain.FileStorage
	timeout   time.Duration
}

func NewReviewService(
	rvRepo domain.ReviewRepository,
	usrRepo domain.UserRepository,
	attrRepo domain.AttractionRepository,
	fileStore domain.FileStorage,
	timeout time.Duration,
) domain.ReviewService {
	return &reviewService{rvRepo, usrRepo, attrRepo, fileStore, timeout}
}

func (s *reviewService) FetchAll(ctx context.Context, args ...interface{}) ([]*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	reviews, err := s.rvRepo.FetchAll(c, args...)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2*len(reviews))

	for _, review := range reviews {
		wg.Add(2)

		go func(review *domain.Review) {
			defer wg.Done()

			userDetail, err := s.usrRepo.FetchOneByArg(c, "user_id", review.UserID)

			if err != nil {
				errs <- err
			}

			review.UserDetail = userDetail
		}(review)

		go func(review *domain.Review) {
			defer wg.Done()
			var attr *domain.Attraction

			attr, err = s.attrRepo.FetchByID(c, "", review.AttractionID)

			if err != nil {
				errs <- err
			}

			review.AttrDetail = *attr
		}(review)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return reviews, nil
}

func (s *reviewService) FetchByAttrID(ctx context.Context, attractionId string, args ...interface{}) ([]*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	reviews, err := s.rvRepo.FetchByAttrID(c, attractionId, args...)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2*len(reviews))

	for _, review := range reviews {
		wg.Add(2)

		go func(review *domain.Review) {
			defer wg.Done()

			userDetail, err := s.usrRepo.FetchOneByArg(c, "user_id", review.UserID)

			if err != nil {
				errs <- err
			}

			review.UserDetail = userDetail
		}(review)

		go func(review *domain.Review) {
			defer wg.Done()
			var attr *domain.Attraction

			attr, err = s.attrRepo.FetchByID(c, "", review.AttractionID)

			if err != nil {
				errs <- err
			}

			review.AttrDetail = *attr
		}(review)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return reviews, nil

}
func (s *reviewService) FetchByID(ctx context.Context, id string, args ...interface{}) (*domain.Review, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	review, err := s.rvRepo.FetchByID(c, id, args...)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errs := make(chan error, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()

		userDetail, err := s.usrRepo.FetchOneByArg(c, "user_id", review.UserID)

		if err != nil {
			errs <- err
		}

		review.UserDetail = userDetail
	}()

	go func() {
		defer wg.Done()
		var attr *domain.Attraction

		attr, err = s.attrRepo.FetchByID(c, "", review.AttractionID)

		if err != nil {
			errs <- err
		}

		review.AttrDetail = *attr
	}()

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return review, nil
}

func (s *reviewService) InsertReview(ctx context.Context, review *domain.ReviewPayload) error {
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

func (s *reviewService) LikeReview(ctx context.Context, reviewId, userId string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	_, err := s.rvRepo.FetchByID(ctx, reviewId)

	if err != nil {
		return domain.ErrNotFound
	}

	err = s.rvRepo.LikeReview(c, reviewId, userId)

	return err
}

func (s *reviewService) UnlikeReview(ctx context.Context, reviewId, userId string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.rvRepo.UnlikeReview(c, reviewId, userId)

	return err
}

func (s *reviewService) DeleteReview(ctx context.Context, reviewId, userId string) error {
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

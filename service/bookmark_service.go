package service

import (
	"context"
	"sync"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/gmaps"
)

type bookmarkService struct {
	bmRepo   domain.BookmarkRepository
	userRepo domain.UserRepository
	attrRepo domain.AttractionRepository
	attrPhotoRepo domain.AttractionPhotoRepository
	attrPriceRepo domain.PriceDetailsRepository
	opHourRepo    domain.OperationHoursRepository
}

func NewBookmarkService(
	bmRepo domain.BookmarkRepository,
	userRepo domain.UserRepository,
	attrRepo domain.AttractionRepository,
	attrPhotoRepo domain.AttractionPhotoRepository,
	attrPriceRepo domain.PriceDetailsRepository,
	opHourRepo    domain.OperationHoursRepository,
) domain.BookmarkService {
	return &bookmarkService{bmRepo, userRepo, attrRepo, attrPhotoRepo, attrPriceRepo, opHourRepo}
}

func (s *bookmarkService) GetBookmarkedByUserID(ctx context.Context, userId string) ([]*domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, 12 * time.Second)
	defer cancel()

	attractions, err := s.bmRepo.GetBookmarkedByUserID(c, userId)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	errs := make(chan error, len(attractions) * 3)

	for _, attraction := range attractions {
		wg.Add(3) 

		go func(attr *domain.Attraction) {
			wg.Done()

			attr.Photos, err = s.attrPhotoRepo.FetchPhotoUrlsByAttrID(c, attr.ID)

			if err != nil {
				errs <- err
			}
		}(attraction)

		go func(attr *domain.Attraction) {
			wg.Done()
			
			attr.Rating, err = gmaps.GetRatings(attr.Name)

			if err != nil {
				errs <- err
			}
		}(attraction)

		go func(attr *domain.Attraction) {
			wg.Done()

			attr.PriceDetails, err = s.attrPriceRepo.FetchByAttID(c, attr.ID)

			if err != nil {
				errs <- err
			}
		}(attraction)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	

	return attractions, err
}

func (s *bookmarkService) InsertBookmark(ctx context.Context, userId, attractionId string) error {
	var wg sync.WaitGroup

	wg.Add(2)

	errs := make(chan error, 2)

	go func() {
		defer wg.Done()

		_, err := s.userRepo.FetchOneByArg(ctx, "user_id", userId)

		if err != nil {
			errs <- err
		}
	}()

	go func() {
		defer wg.Done()

		_, err := s.attrRepo.FetchByID(ctx, attractionId)

		if err != nil {
			errs <- err
		}
	}()

	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		if err != nil {
			return err
		}
	}

	return s.bmRepo.InsertBookmark(ctx, userId, attractionId)
}

func(s *bookmarkService) RemoveBookmark(ctx context.Context, userId, attractionId string) error {
	c, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	return s.bmRepo.RemoveBookmark(c, userId, attractionId)
}
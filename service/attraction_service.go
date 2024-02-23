package service

import (
	"context"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

const ATTRACTION_CLOUD_STORE_DIR = "attractions"

type attractionService struct {
	attrRepo domain.AttractionRepository
	attrPhotoRepo domain.AttractionPhotoRepository
	fileStore domain.FileStorage
	timeout time.Duration 
}

func NewAttractionSerivce(
	attrRepo domain.AttractionRepository, 
	attrPhotoRepo domain.AttractionPhotoRepository,
	fileStore domain.FileStorage, 
	timeout time.Duration,
) domain.AttractionService {
	return &attractionService{attrRepo, attrPhotoRepo, fileStore, timeout}
}

func(s *attractionService) FetchAll(ctx context.Context) ([]domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attractions, err := s.attrRepo.FetchAll(c)

	for _, attraction := range attractions {
		attraction.Photos, err = s.attrPhotoRepo.FetchPhotoUrlsByAttrID(c, attraction.ID)

		if err != nil {
			return nil, err 
		}
	}

	return attractions, err 
}

func(s *attractionService) FetchByID(ctx context.Context, id string) (domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attraction, err := s.attrRepo.FetchByID(c, id)
	
	if err != nil {
		return domain.Attraction{}, err 
	}

	attraction.Photos, err = s.attrPhotoRepo.FetchPhotoUrlsByAttrID(ctx, id)

	return attraction, err 
}

func(s *attractionService) InsertAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)	
	defer cancel()

	attraction.Default()

	err := s.attrRepo.InsertAttraction(c, attraction)

	return err 
}

func(s *attractionService) UpdateAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)	
	defer cancel()

	err := s.attrRepo.UpdateAttraction(c, attraction) 

	return err 
}

func(s *attractionService) UploadPhotoByAttID(ctx context.Context, attrPhoto *domain.AttractionPhoto) error  {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	_, err := s.attrRepo.FetchByID(c, attrPhoto.AttractionID)

	if err != nil {
		return err 
	}

	attrPhoto.PhotoUrl, err = s.fileStore.UploadFile(ATTRACTION_CLOUD_STORE_DIR, attrPhoto.PhotoFile)

	if err != nil {
		return err
	}

	err = s.attrPhotoRepo.InsertPhotoUrl(ctx, attrPhoto)

	return err 
}

func(s *attractionService) DeleteAttraction(ctx context.Context, id string) error  {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	err := s.attrRepo.DeleteAttraction(c, id)

	return err 
}
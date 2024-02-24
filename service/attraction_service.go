package service

import (
	"context"
	"sync"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

const ATTRACTION_CLOUD_STORE_DIR = "attractions"

type attractionService struct {
	attrRepo      domain.AttractionRepository
	attrPhotoRepo domain.AttractionPhotoRepository
	opHourRepo    domain.OperationHoursRepository
	fileStore     domain.FileStorage
	timeout       time.Duration
}

func NewAttractionSerivce(
	attrRepo domain.AttractionRepository,
	attrPhotoRepo domain.AttractionPhotoRepository,
	opHourRepo domain.OperationHoursRepository,
	fileStore domain.FileStorage,
	timeout time.Duration,
) domain.AttractionService {
	return &attractionService{attrRepo, attrPhotoRepo, opHourRepo, fileStore, timeout}
}

func (s *attractionService) FetchAll(ctx context.Context) ([]*domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attractions, err := s.attrRepo.FetchAll(c)

	for _, attraction := range attractions {

		attraction.Photos, err = s.attrPhotoRepo.FetchPhotoUrlsByAttrID(c, attraction.ID)

		if err != nil {
			return nil, err
		}

		attraction.OperationHours, err = s.opHourRepo.FetchByAttID(c, attraction.ID)

		if err != nil {
			return nil, err
		}
	}

	return attractions, err
}

func (s *attractionService) FetchByID(ctx context.Context, id string) (*domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attraction, err := s.attrRepo.FetchByID(c, id)

	if err != nil {
		return nil, err
	}

	attraction.Photos, err = s.attrPhotoRepo.FetchPhotoUrlsByAttrID(ctx, id)

	if err != nil {
		return nil, err
	}

	attraction.OperationHours, err = s.opHourRepo.FetchByAttID(c, id)

	return attraction, err
}

func (s *attractionService) InsertAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attraction.Default()

	errChan := make(chan error, len(attraction.OpeningHours))

	var wg sync.WaitGroup

	for _, opHour := range attraction.OperationHours {

		opHour.Default(attraction.ID)
		wg.Add(1)

		go func(opHour domain.OperationHours) {
			defer wg.Done()

			err := s.opHourRepo.InsertWithAttID(ctx, &opHour)

			if err != nil {
				errChan <- err
			}
		}(*opHour)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	err := s.attrRepo.InsertAttraction(c, attraction)

	return err
}

func (s *attractionService) UpdateAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.attrRepo.UpdateAttraction(c, attraction)

	return err
}

func (s *attractionService) UploadPhotoByAttID(ctx context.Context, attrPhoto *domain.AttractionPhotoPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	_, err := s.attrRepo.FetchByID(c, attrPhoto.AttractionID)

	if err != nil {
		return err
	}

	errChan := make(chan error, len(attrPhoto.PhotoFiles) * 2)

	var wg sync.WaitGroup

	for _, file := range attrPhoto.PhotoFiles {
		attrPh := domain.AttractionPhoto{}
		attrPh.PhotoFile = file 
		attrPh.AttractionID = attrPhoto.AttractionID

		wg.Add(1)

		go func(attrPh domain.AttractionPhoto) {
			defer wg.Done()

			attrPh.PhotoUrl, err = s.fileStore.UploadFile(ATTRACTION_CLOUD_STORE_DIR, attrPh.PhotoFile)

			if err != nil {
				errChan <- err 
			}

			err = s.attrPhotoRepo.InsertPhotoUrl(ctx, &attrPh)

			if err != nil {
				errChan <- err 
			}

		}(attrPh)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err 
		}
	}

	return nil
}

func (s *attractionService) DeleteAttraction(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	err := s.attrRepo.DeleteAttraction(c, id)

	return err
}

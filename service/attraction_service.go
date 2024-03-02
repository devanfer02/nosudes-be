package service

import (
	"context"
	"sync"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/gmaps"
	"github.com/devanfer02/nosudes-be/utils/weather"
)

const ATTRACTION_CLOUD_STORE_DIR = "attractions"

type attractionService struct {
	attrRepo      domain.AttractionRepository
	attrPhotoRepo domain.AttractionPhotoRepository
	attrPriceRepo domain.PriceDetailsRepository
	opHourRepo    domain.OperationHoursRepository
	fileStore     domain.FileStorage
	timeout       time.Duration
}

func NewAttractionSerivce(
	attrRepo domain.AttractionRepository,
	attrPhotoRepo domain.AttractionPhotoRepository,
	attrPriceRepo domain.PriceDetailsRepository,
	opHourRepo domain.OperationHoursRepository,
	fileStore domain.FileStorage,
	timeout time.Duration,
) domain.AttractionService {
	return &attractionService{attrRepo, attrPhotoRepo, attrPriceRepo, opHourRepo, fileStore, timeout}
}

func (s *attractionService) FetchAll(ctx context.Context, query *domain.LocQuery) ([]*domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attractions, err := s.attrRepo.FetchAll(c)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	errChan := make(chan error, len(attractions) * 5)

	for _, attraction := range attractions {
		wg.Add(1)
		go func(attr *domain.Attraction){
			defer wg.Done()
			s.fetch(c, attr, errChan, query)
		}(attraction)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return attractions, err
		}
	}

	return attractions, err
}

func (s *attractionService) FetchByID(ctx context.Context, id string, query *domain.LocQuery) (*domain.Attraction, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attraction, err := s.attrRepo.FetchByID(c, id)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	errChan := make(chan error, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.fetch(c, attraction, errChan, query)
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return attraction, err
		}
	}

	return attraction, err
}


func (s *attractionService) InsertAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attraction.Default()

	err := s.attrRepo.InsertAttraction(c, attraction)

	if err != nil {
		return err
	}

	// concurrency to insert multipler opening hours
	errChan := make(chan error, len(attraction.OpeningHours))

	var wgOp sync.WaitGroup

	for _, opHour := range attraction.OperationHours {

		opHour.Default(attraction.ID)
		wgOp.Add(1)

		go func(opHour domain.OperationHours) {
			defer wgOp.Done()

			err := s.opHourRepo.InsertWithAttID(ctx, &opHour)

			if err != nil {
				errChan <- err
			}
		}(*opHour)
	}

	go func() {
		wgOp.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// concurrency to insert mutlipler price details

	errChan = make(chan error, len(attraction.PriceDetails))

	var wgPr sync.WaitGroup

	for _, prd := range attraction.PriceDetails {
		wgPr.Add(1)

		prd.AttractionID = attraction.ID

		go func(price domain.PriceDetails) {
			defer wgPr.Done()

			err := s.attrPriceRepo.InsertWithAttID(ctx, &price)

			if err != nil {
				errChan <- err
			}
		}(*prd)
	}

	go func() {
		wgPr.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
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

	errChan := make(chan error, len(attrPhoto.PhotoFiles)*2)

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

func (s *attractionService) fetch(
	c context.Context, 
	attr *domain.Attraction, 
	errChan chan error, 
	query *domain.LocQuery,
) {
	var wg sync.WaitGroup 

	wg.Add(3)

	go func() {
		defer wg.Done()
		photos, err := s.attrPhotoRepo.FetchPhotoUrlsByAttrID(c, attr.ID)

		if err != nil {
			errChan <- err
			return
		}
		
		attr.Photos = photos
	}()

	go func() {
		defer wg.Done()
		detail, err := gmaps.GetRatings(attr.Name)

		if err != nil {
			errChan <- err
			return 
		}

		attr.MapsDetail = &detail
	}()

	go func() {
		defer wg.Done()
		priceDetails, err := s.attrPriceRepo.FetchByAttID(c, attr.ID)

		if err != nil {
			errChan <- err
			return
		}

		attr.PriceDetails = priceDetails
	}()

	wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		
		weatherInfo, err := weather.GetWeatherInfo(
			attr.MapsDetail.GeometryInfo.Loc.Lat, 
			attr.MapsDetail.GeometryInfo.Loc.Lng, 
		)

		if err != nil {
			errChan <- err 
			return 
		}

		attr.WeatherInfo = weatherInfo
	}()

	wg.Wait()

	if query == nil {
		return 
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		dst, err := gmaps.GetDistance(attr, query)

		if err != nil {
			errChan <- err
		}

		attr.DistanceValue = dst.DistanceValue
	}()

	wg.Wait()
}
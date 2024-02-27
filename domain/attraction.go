package domain

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
)

type Attraction struct {
	ID             string            `json:"id" db:"attraction_id"`
	Name           string            `json:"name" db:"name"`
	Category       string            `json:"category" db:"category"`
	Description    string            `json:"description" db:"description"`
	OpeningHours   string            `json:"opening_hours" db:"opening_hours"`
	MapsEmbedUrl   string            `json:"maps_embed_url" db:"maps_embed_url"`
	Location       string            `json:"location" db:"location"`
	OperationHours []OperationHours  `json:"operation_hours,omitempty" db:"-"`
	PriceDetails   []PriceDetails    `json:"price_details,omitempty" db:"-"`
	Photos         []AttractionPhoto `json:"attraction_photos,omitempty" db:"-"`
	Rating         Ratings           `json:"ratings"`
	CreatedAt      time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at" db:"updated_at"`
}

type AttractionPayload struct {
	ID             string            `json:"id" db:"attraction_id"`
	Name           string            `json:"name" db:"name" binding:"required"`
	CategoryID     int               `json:"category_id" db:"category_id" binding:"required"`
	Description    string            `json:"description" db:"description" binding:"required"`
	OpeningHours   string            `json:"opening_hours" db:"opening_hours" binding:"required"`
	MapsEmbedUrl   string            `json:"maps_embed_url" db:"maps_embed_url"`
	Location       string            `json:"location" db:"location" binding:"required"`
	OperationHours []*OperationHours `json:"operation_hours,omitempty" db:"-"`
	PriceDetails   []*PriceDetails   `json:"price_details,omitempty" db:"-"`
}

type GmapsRef struct {
	Results []Ratings `json:"results"`
}

type Ratings struct {
	Address      string  `json:"formatted_address"`
	Rating       float64 `json:"rating"`
	TotalRatings int64   `json:"user_ratings_total"`
}

type PriceDetails struct {
	AttractionID string `json:"-" db:"attraction_id"`
	Price        int64  `json:"price" db:"price"`
	DayType      string `json:"day_type,omitempty" db:"day_type"`
	AgeGroup     string `json:"age_group,omitempty" db:"age_group"`
}

type AttractionPhoto struct {
	AttractionID string                `json:"-" db:"attraction_id" form:"-"`
	PhotoUrl     string                `json:"photo_url" db:"photo_url" form:"-"`
	PhotoFile    *multipart.FileHeader `json:"-"`
}

type AttractionPhotoPayload struct {
	AttractionID string                  `json:"-"`
	PhotoFiles   []*multipart.FileHeader `form:"photos" binding:"required"`
}

type OperationHours struct {
	OpHoursID    string `json:"id" db:"op_hour_id"`
	AttractionID string `json:"attraction_id" db:"attraction_id"`
	Day          string `json:"day" db:"day" binding:"required"`
	DayIndex     int    `json:"-" db:"day_index"`
	Timespan     string `json:"timespan" db:"timespan" binding:"required"`
}

type AttractionRepository interface {
	FetchAll(ctx context.Context) ([]*Attraction, error)
	FetchByID(ctx context.Context, id string) (*Attraction, error)
	InsertAttraction(ctx context.Context, attraction *AttractionPayload) error
	UpdateAttraction(ctx context.Context, attraction *AttractionPayload) error
	DeleteAttraction(ctx context.Context, id string) error
}

type AttractionPhotoRepository interface {
	FetchPhotoUrlsByAttrID(ctx context.Context, attractionId string) ([]AttractionPhoto, error)
	InsertPhotoUrl(ctx context.Context, attr *AttractionPhoto) error
}

type OperationHoursRepository interface {
	FetchByAttID(ctx context.Context, attractionId string) ([]OperationHours, error)
	InsertWithAttID(ctx context.Context, ophour *OperationHours) error
	UpdateByID(ctx context.Context, ophour *OperationHours) error
	DeleteByID(ctx context.Context, id string) error
}

type PriceDetailsRepository interface {
	FetchByAttID(ctx context.Context, attractionId string) ([]PriceDetails, error)
	InsertWithAttID(ctx context.Context, price *PriceDetails) error
}

type AttractionService interface {
	FetchAll(ctx context.Context) ([]*Attraction, error)
	FetchByID(ctx context.Context, id string) (*Attraction, error)
	InsertAttraction(ctx context.Context, attraction *AttractionPayload) error
	UpdateAttraction(ctx context.Context, attraction *AttractionPayload) error
	UploadPhotoByAttID(ctx context.Context, attPhoto *AttractionPhotoPayload) error
	DeleteAttraction(ctx context.Context, id string) error
}

func (a *AttractionPayload) Default() {
	a.MapsEmbedUrl = fmt.Sprintf("https://www.google.com/maps/embed/v1/place?key=%s&q=%s", env.ProcEnv.MapsAPIKey, a.Name)
	a.MapsEmbedUrl = strings.ReplaceAll(a.MapsEmbedUrl, " ", "+")
	a.ID = uuid.New().String()
}

func (o *OperationHours) Default(attractionId string) {
	o.OpHoursID = uuid.New().String()
	o.AttractionID = attractionId
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for idx, day := range days {
		if o.Day == day {
			o.DayIndex = idx + 1
			break
		}
	}

	if o.DayIndex == 0 {
		logger.ErrLog(layers.Domain, "weird behaviour, day not match anything", fmt.Errorf("day %s not match in array", o.Day))
	}

}

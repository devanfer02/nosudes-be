package domain

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
)

type Attraction struct {
	ID             string            `json:"id" db:"attraction_id"`
	Name           string            `json:"name" db:"name"`
	Category       string            `json:"category" db:"category"`
	Description    string            `json:"description" db:"description"`
	OpeningHours   string            `json:"opening_hours" db:"opening_hours"`
	MapsEmbedUrl   string            `json:"maps_embed_url" db:"maps_embed_url"`
	OperationHours []OperationHours  `json:"operation_hours,omitempty" db:"-"`
	PriceDetails   []PriceDetails    `json:"price_details,omitempty" db:"-"`
	Photos         []AttractionPhoto `json:"attraction_photos,omitempty" db:"-"`
	CreatedAt      time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at" db:"updated_at"`
}

type AttractionPayload struct {
	ID           string `json:"id" db:"attraction_id"`
	Name         string `json:"name" db:"name" binding:"required"`
	CategoryID   int    `json:"category_id" db:"category_id" binding:"required"`
	Description  string `json:"description" db:"description" binding:"required"`
	OpeningHours string `json:"opening_hours" db:"opening_hours" binding:"required"`
	MapsEmbedUrl string `json:"maps_embed_url" db:"maps_embed_url"`
}

type PriceDetails struct {
	AttractionID string `json:"-" db:"attraction_id"`
	Price        string `json:"price" db:"price"`
}

type AttractionPhoto struct {
	AttractionID string                `json:"-" db:"attraction_id" form:"-"`
	PhotoUrl     string                `json:"photo_url" db:"photo_url" form:"-"`
	PhotoFile    *multipart.FileHeader `form:"photo" json:"-" db:"-" binding:"required"`
}

type AttractionRepository interface {
	FetchAll(ctx context.Context) ([]Attraction, error)
	FetchByID(ctx context.Context, id string) (Attraction, error)
	InsertAttraction(ctx context.Context, attraction *AttractionPayload) error
	UpdateAttraction(ctx context.Context, attraction *AttractionPayload) error
	DeleteAttraction(ctx context.Context, id string) error
}

type AttractionPhotoRepository interface {
	FetchPhotoUrlsByAttrID(ctx context.Context, attractionId string) ([]AttractionPhoto, error)
	InsertPhotoUrl(ctx context.Context, attr *AttractionPhoto) error
}

type AttractionService interface {
	FetchAll(ctx context.Context) ([]Attraction, error)
	FetchByID(ctx context.Context, id string) (Attraction, error)
	InsertAttraction(ctx context.Context, attraction *AttractionPayload) error
	UpdateAttraction(ctx context.Context, attraction *AttractionPayload) error
	UploadPhotoByAttID(ctx context.Context, attPhoto *AttractionPhoto) error
	DeleteAttraction(ctx context.Context, id string) error
}

func (a *AttractionPayload) Default() {
	a.MapsEmbedUrl = fmt.Sprintf("https://www.google.com/maps/embed/v1/place?key=%s&q=%s", env.ProcEnv.MapsAPIKey, a.Name)
	a.MapsEmbedUrl = strings.ReplaceAll(a.MapsEmbedUrl, " ", "+")
	a.ID = uuid.New().String()
}

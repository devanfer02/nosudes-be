package domain

import (
	"context"
)

type OperationHours struct {
	OpHoursID	 string `json:"id" db:"op_hour_id"`
	AttractionID string `json:"attraction_id" db:"attraction_id"`
	Day          string `json:"day" db:"day"`
	Timespan     string `json:"timespan" db:"timespan"`
}

type OperationHoursRepository interface {
	FetchByAttID(ctx context.Context) ([]OperationHours, error)
	InsertWithAttID(ctx context.Context, ophour *OperationHours) error 
	UpdateByID(ctx context.Context, ophour *OperationHours) error 
	DeleteByID(ctx context.Context, id string) error 
}

type OperationHoursService interface {
	FetchByAttID(ctx context.Context) ([]OperationHours, error)
	InsertWithAttID(ctx context.Context, ophour *OperationHours) error 
	UpdateByID(ctx context.Context, ophour *OperationHours) error 
	DeleteByID(ctx context.Context, id string) error 
}
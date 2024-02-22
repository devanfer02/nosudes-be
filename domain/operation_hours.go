package domain

import (
	"context"
)

type OperationHours struct {
	AttractionID string `json:"attraction_id" db:"attraction_id"`
	Day          string `json:"day" db:"day"`
	Timespan     string `json:"timespan" db:"timespan"`
}

type OperationHoursRepository interface {
	FetchByAttID(ctx context.Context) (OperationHours, error)
	InsertWithAttID(ctx context.Context, ophour *OperationHours) error 
	UpdateWithAttID(ctx context.Context, ophour *OperationHours) error 
}
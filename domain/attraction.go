package domain

import "context"

type Attraction struct {
	ID             string           `json:"id" db:"attraction_id"`
	Name           string           `json:"name" db:"name"`
	Category       string           `json:"category" db:"category"`
	Description    string           `json:"description" db:"description"`
	OpeningHours   string           `json:"opening_hours" db:"opening_hours"`
	OperationHours []OperationHours `json:"operation_horus" db:"-"`
	PriceDetails   []string         `json:"price_details" db:"-"`
}

type AttractionPrice struct {
}

type AttractionRepository interface {
	FetchAll(ctx context.Context)
}

type AttractionService interface {
}

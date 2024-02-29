package repository

import (
	"context"

	"github.com/devanfer02/nosudes-be/domain"

	"github.com/jmoiron/sqlx"
)

type mysqlAttractionPricesRepository struct {
	Conn *sqlx.DB
}

func NewMysqlAttractionPricesRepository(conn *sqlx.DB) domain.PriceDetailsRepository {
	return &mysqlAttractionPricesRepository{conn}
}

func(m *mysqlAttractionPricesRepository) FetchByAttID(ctx context.Context, attractionId string) ([]domain.PriceDetails, error) {
	query := `SELECT * FROM attraction_ticket_prices WHERE attraction_id = ?`

	prices := make([]domain.PriceDetails, 0)

	err := m.Conn.SelectContext(ctx, &prices, query, attractionId)

	if err != nil {
		return nil, err 
	}

	return prices, nil
}

func(m *mysqlAttractionPricesRepository) InsertWithAttID(ctx context.Context, price *domain.PriceDetails) error  {
	query := `INSERT INTO attraction_ticket_prices 
		(attraction_id, price, day_type, age_group, park_type)
		VALUES (?, ?, ?, ?, ?)`

	return execStatement(m.Conn, ctx, query, price.AttractionID, price.Price, price.DayType, price.AgeGroup, price.ParkType)
}

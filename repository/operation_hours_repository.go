package repository

import (
	"context"

	"github.com/devanfer02/nosudes-be/domain"

	"github.com/jmoiron/sqlx"
)

type mysqlOpHoursRepository struct {
	Conn *sqlx.DB
}

func NewMysqlOpHoursRepository(conn *sqlx.DB) domain.OperationHoursRepository {
	return	&mysqlOpHoursRepository{conn}
}

func(m *mysqlOpHoursRepository) FetchByAttID(ctx context.Context) ([]domain.OperationHours, error) {
	query := `SELECT * FROM `
}

func(m *mysqlOpHoursRepository) InsertWithAttID(ctx context.Context, ophour *domain.OperationHours) error  {

}

func(m *mysqlOpHoursRepository) UpdateByID(ctx context.Context, ophour *domain.OperationHours) error  {

}

func(m *mysqlOpHoursRepository) DeleteByID(ctx context.Context, id string) error  {

}


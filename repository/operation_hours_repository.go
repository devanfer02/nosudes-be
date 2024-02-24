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

func(m *mysqlOpHoursRepository) FetchByAttID(ctx context.Context, attractionId string) ([]domain.OperationHours, error) {
	query := `SELECT * FROM operation_hours WHERE attraction_id = ? ORDER BY day_index ASC`

	opHours := make([]domain.OperationHours, 0)

	err := m.Conn.SelectContext(ctx, &opHours, query, attractionId)

	if err != nil {
		return nil, err
	}

	return opHours, nil
}

func(m *mysqlOpHoursRepository) InsertWithAttID(ctx context.Context, ophour *domain.OperationHours) error  {
	query := `INSERT INTO operation_hours (op_hour_id, attraction_id, day, day_index, timespan) VALUES (?, ?, ?, ?, ?)`

	return execStatement(m.Conn, ctx, query, ophour.OpHoursID, ophour.AttractionID, ophour.Day, ophour.DayIndex, ophour.Timespan)
}

func(m *mysqlOpHoursRepository) UpdateByID(ctx context.Context, ophour *domain.OperationHours) error  {
	query := `UPDATE operation_hours SET day = ?, timespan = ?`

	return execStatement(m.Conn, ctx, query, ophour.Day, ophour.Timespan)
}

func(m *mysqlOpHoursRepository) DeleteByID(ctx context.Context, id string) error  {
	query := `DELETE FROM operation_hours WHERE op_hour_id = ?`

	return execStatement(m.Conn, ctx, query, id)
}


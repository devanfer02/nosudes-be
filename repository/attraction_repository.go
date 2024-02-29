package repository

import (
	"context"
	"database/sql"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
	helpers "github.com/devanfer02/nosudes-be/utils"
	
	"github.com/jmoiron/sqlx"
)

type mysqlAttractionRepository struct {
	Conn *sqlx.DB
}

func NewMysqlAttractionRepository(conn *sqlx.DB) domain.AttractionRepository {
	return &mysqlAttractionRepository{conn}
}

func(m *mysqlAttractionRepository) FetchAll(ctx context.Context) ([]*domain.Attraction, error) {
	query := `SELECT 
		a.attraction_id AS attraction_id, 
		a.name, ac.category_name AS category,
		description, opening_hours, maps_embed_url
		FROM attractions a JOIN attraction_categories ac
		ON a.category_id = ac.category_id`

	attractions := make([]*domain.Attraction, 0)

	err := m.Conn.SelectContext(ctx, &attractions, query)

	if err != nil {
		logger.ErrLog(layers.Repository, "error fetching attractions data", err)
		return nil, err 
	}

	return attractions, nil 
}

func(m *mysqlAttractionRepository) FetchByID(ctx context.Context, id string) (*domain.Attraction, error) {
	query := `SELECT 
		a.attraction_id AS attraction_id, 
		a.name, ac.category_name AS category,
		description, opening_hours, maps_embed_url
		FROM attractions a JOIN attraction_categories ac
		ON a.category_id = ac.category_id
		WHERE attraction_id = ?`

	attraction := &domain.Attraction{}

	err := m.Conn.GetContext(ctx, attraction, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		logger.ErrLog(layers.Repository, "failed to exec query", err)
		return nil, domain.ErrInternalServer
	}

	return attraction, nil
}

func(m *mysqlAttractionRepository) InsertAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	query := `INSERT INTO 
	attractions (attraction_id, name, category_id, description, opening_hours, maps_embed_url, location, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	currTime := helpers.CurrentTime()

	return execStatement(
		m.Conn, ctx, query, 
		attraction.ID, 
		attraction.Name, 
		attraction.CategoryID,
		attraction.Description,
		attraction.OpeningHours, 
		attraction.MapsEmbedUrl, 
		attraction.Location,
		currTime, currTime,
	)
}

func(m *mysqlAttractionRepository) UpdateAttraction(ctx context.Context, attraction *domain.AttractionPayload) error {
	query := `UPDATE attractions SET name = ?, category_id = ?, description = ?, opening_hours = ?, maps_embed_url = ?, updated_at WHERE attraction_id = ?`

	currTime := helpers.CurrentTime()

	return execStatement(
		m.Conn, ctx, query,
		attraction.Name,
		attraction.CategoryID,
		attraction.Description,
		attraction.OpeningHours,
		attraction.MapsEmbedUrl,
		currTime,
		attraction.ID, 
	)
}

func(m *mysqlAttractionRepository) DeleteAttraction(ctx context.Context, id string) error  {
	query := `DELETE FROM attractions WHERE attraction_id = ?`

	return execStatement(
		m.Conn, ctx, query,
		id,
	)
}


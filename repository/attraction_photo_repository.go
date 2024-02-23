package repository

import (
	"context"

	"github.com/devanfer02/nosudes-be/domain"

	"github.com/jmoiron/sqlx"
)

type mysqlAttractionPhotoReqository struct {
	Conn *sqlx.DB
}

func NewMysqlAttractionPhotoRepository(conn *sqlx.DB) domain.AttractionPhotoRepository {
	return &mysqlAttractionPhotoReqository{conn}
}

func (m *mysqlAttractionPhotoReqository) FetchPhotoUrlsByAttrID(ctx context.Context, attractionId string) ([]domain.AttractionPhoto, error) {
	query := `SELECT * FROM attraction_photos WHERE attraction_id = ?`

	photoUrls := make([]domain.AttractionPhoto, 0)

	err := m.Conn.SelectContext(ctx, &photoUrls, query, attractionId)

	if err != nil {
		return nil, err
	}

	return photoUrls, nil
}

func (m *mysqlAttractionPhotoReqository) InsertPhotoUrl(ctx context.Context, attr *domain.AttractionPhoto) error {
	query := `INSERT INTO attraction_photos (attraction_id, photo_url) VALUES (?, ?)`

	return execStatement(m.Conn, ctx, query, attr.AttractionID, attr.PhotoUrl)
}

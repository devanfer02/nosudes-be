package repository

import (
	"context"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
	
	"github.com/jmoiron/sqlx"	
)

type mysqlBookmarkRepository struct {
	Conn *sqlx.DB
}

func NewMysqlBookmarkRepository(conn *sqlx.DB) domain.BookmarkRepository {
	return &mysqlBookmarkRepository{conn}
}

func(m *mysqlBookmarkRepository) GetBookmarkedByUserID(ctx context.Context, userId string) ([]*domain.Attraction, error) {
	query := `SELECT a.* FROM attractions a JOIN bookmarks b ON b.attraction_id = ? WHERE b.user_id = ?`

	attractions := make([]*domain.Attraction, 0)

	err := m.Conn.SelectContext(ctx, &attractions, query, userId)

	if err != nil {
		logger.ErrLog(layers.Repository, "error fetching attractions data", err)
		return nil, err 
	}

	return attractions, nil 
}

func(m *mysqlBookmarkRepository) InsertBookmark(ctx context.Context, userId, attractionId string) error {
	query := `INSERT INTO bookmarks (attraction_id, user_id) VALUES (?, ?)`

	return execStatement(m.Conn, ctx, query, attractionId, userId)
}

func(m *mysqlBookmarkRepository) RemoveBookmark(ctx context.Context, userId, attractionId string) error {
	query := `DELETE FROM bookmarks WHERE attraction_id = ? AND user_id = ?`

	return execStatement(m.Conn, ctx, query, attractionId, userId)
}
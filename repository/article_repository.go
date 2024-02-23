package repository

import (
	"context"
	"database/sql"

	"github.com/devanfer02/nosudes-be/domain"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
	
	"github.com/jmoiron/sqlx"
)

type mysqlArticleRepository struct {
	Conn *sqlx.DB
}

func NewMysqlArticleRepository(conn *sqlx.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{conn}
}

func(m *mysqlArticleRepository) FetchAll(ctx context.Context) ([]domain.Article, error)  {
	query := "SELECT * FROM articles"

	articles := make([]domain.Article, 0)

	err := m.Conn.SelectContext(ctx, &articles, query)

	if err != nil {
		logger.ErrLog(layers.Repository, "failed to exec query", err)
		return nil, domain.ErrInternalServer
	}

	return articles, nil
}

func(m *mysqlArticleRepository) FetchByID(ctx context.Context, id string) (domain.Article, error) {
	query := "SELECT * FROM articles WHERE article_id = ?"

	article := domain.Article{}

	err := m.Conn.GetContext(ctx, &article, query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Article{}, domain.ErrNotFound
		}
		logger.ErrLog(layers.Repository, "failed to exec query", err)
		return domain.Article{}, domain.ErrInternalServer
	}

	return article, nil
}

func(m *mysqlArticleRepository) InsertArticle(ctx context.Context, article *domain.Article) error  {
	query := `INSERT INTO articles 
		(article_id, title, summary, description, photo)
		VALUES
		(?, ?, ?, ?, ?)`

	
	return execStatement(m.Conn, ctx, query, article.ID, article.Title, article.Summary, article.Description, article.Photo)
}

func(m *mysqlArticleRepository) UpdateArticle(ctx context.Context, article *domain.Article) error {
	query := `UPDATE articles SET title = ?, summary = ?, description = ?, photo = ? WHERE article_id = ?`

	return execStatement(m.Conn, ctx, query, article.Title, article.Summary, article.Description, article.Photo, article.ID)
}

func(m *mysqlArticleRepository) DeleteArticle(ctx context.Context, id string) error {
	query := `DELETE FROM articles WHERE article_id = ?`

	return execStatement(m.Conn, ctx, query, id)
}
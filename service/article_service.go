package service

import (
	"context"
	"time"

	"github.com/devanfer02/nosudes-be/domain"
)

const FIREBASE_STORE_DIR = "articles"

type articleService struct {
	artRepo domain.ArticleRepository
	fileStore domain.FileStorage
	timeout time.Duration
}

func NewArticleService(artRepo domain.ArticleRepository, fileStore domain.FileStorage, timeout time.Duration) domain.ArticleService {
	return &articleService{artRepo, fileStore, timeout}
}

func(s *articleService) FetchAll(ctx context.Context) ([]domain.Article, error)  {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	articles, err := s.artRepo.FetchAll(c)

	return articles, err
}

func(s *articleService) FetchByID(ctx context.Context, id string) (domain.Article, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	article, err := s.artRepo.FetchByID(c, id)

	return article, err
}

func(s *articleService) InsertArticle(ctx context.Context, article *domain.ArticlePayload) error  {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	url, err := s.fileStore.UploadFile(FIREBASE_STORE_DIR, article.PhotoFile)

	if err != nil {
		return err 
	}

	err = s.artRepo.InsertArticle(c, article.Convert(url))

	return err
}

func(s *articleService) UpdateArticle(ctx context.Context, article *domain.ArticlePayload) error  {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	url, err := s.fileStore.UploadFile(FIREBASE_STORE_DIR, article.PhotoFile)

	if err != nil {
		return err 
	}

	err = s.artRepo.UpdateArticle(c, article.Convert(url))

	return err 
}

func(s *articleService) DeleteArticle(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	err := s.artRepo.DeleteArticle(c, id)

	return err 
}

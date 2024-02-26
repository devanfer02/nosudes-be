package domain

import (
	"context"
)

type BookmarkRepository interface {
	GetBookmarkedByUserID(ctx context.Context, userId string) ([]*Attraction, error)
	InsertBookmark(ctx context.Context, userId, attractionId string) error
	RemoveBookmark(ctx context.Context, userId, attractionId string) error
}

type BookmarkService interface {
	GetBookmarkedByUserID(ctx context.Context, userId string) ([]*Attraction, error)
	InsertBookmark(ctx context.Context, userId, attractionId string) error
	RemoveBookmark(ctx context.Context, userId, attractionId string) error
}
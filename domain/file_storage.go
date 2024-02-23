package domain

import (
	"context"
	"mime/multipart"
)

type FileStorage interface {
	UploadFile(dir string, file *multipart.FileHeader) (string, error)
	DeleteFile(ctx context.Context, downloadUrl string) error
}

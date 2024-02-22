package domain

import (
	"mime/multipart"
)

type FileStorage interface {
	UploadFile(dir string, file *multipart.FileHeader) (string, error) 	
}
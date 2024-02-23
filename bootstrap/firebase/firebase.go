package firebase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	filepath "path/filepath"
	"strings"

	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/domain"
	helpers "github.com/devanfer02/nosudes-be/utils"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"

	"github.com/google/uuid"

	cstorage "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	fstorage "firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

type FirebaseStorage struct {
	client *fstorage.Client
}

func NewFirebaseStorage() domain.FileStorage {
	configStore := &firebase.Config{
		StorageBucket: env.ProcEnv.FirebaseBucket,
	}
	opt := option.WithCredentialsFile(env.ProcEnv.FirebaseConf)

	app, err := firebase.NewApp(context.Background(), configStore, opt)

	if err != nil {
		logger.FatalLog(layers.Firebase, "error initializing firebase", err)
	}

	client, err := app.Storage(context.Background())

	if err != nil {
		logger.FatalLog(layers.Firebase, "error getting storage", err)
	}

	return &FirebaseStorage{client}
}

func (f *FirebaseStorage) UploadFile(dir string, file *multipart.FileHeader) (string, error) {

	ctx := context.Background()

	currTime := strings.Replace(helpers.CurrentTime(), " ", "_", 5)

	bucket, err := f.client.DefaultBucket()

	if err != nil {
		logger.ErrLog(layers.Firebase, "error getting bucket", err)
		return "", err
	}

	src, err := file.Open()

	if err != nil {
		logger.ErrLog(layers.Firebase, "error opening file", err)
		return "", err
	}
	defer src.Close()

	uuid := uuid.New().String()

	filename := dir + "/" + currTime + "_" + uuid + filepath.Ext(file.Filename)

	obj := bucket.Object(filename)
	wr := obj.NewWriter(ctx)

	wr.ObjectAttrs.Metadata = map[string]string{
		"firebaseStorageDownloadTokens": uuid,
	}

	if _, err := io.Copy(wr, src); err != nil {
		logger.ErrLog(layers.Firebase, "error copying file to bucket", err)
		return "", err
	}

	if err := wr.Close(); err != nil {
		return "", err
	}

	url, err := f.CreateDownloadUrl(ctx, obj)

	if err != nil {
		return "", err
	}

	return url, nil
}

func (f *FirebaseStorage) CreateDownloadUrl(ctx context.Context, obj *cstorage.ObjectHandle) (string, error) {
	attrs, err := obj.Attrs(ctx)

	if err != nil {
		return "", err
	}

	token, ok := attrs.Metadata["firebaseStorageDownloadTokens"]

	if !ok {
		return "", fmt.Errorf("file token not found")
	}

	parsed := url.QueryEscape(obj.ObjectName())

	return fmt.Sprintf(
		"https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s",
		obj.BucketName(),
		parsed,
		token,
	), nil
}

func (f *FirebaseStorage) DeleteFile(ctx context.Context, downloadUrl string) error {
	parsedUrl, err := url.Parse(downloadUrl)

	if err != nil {
		logger.ErrLog(layers.Firebase, "error parsing url", err)
		return err
	}

	objectPath := strings.TrimLeft(parsedUrl.Path, "/")

	fmt.Println(objectPath)

	bucket, err := f.client.DefaultBucket()

	if err != nil {
		return err
	}

	object := bucket.Object(objectPath)

	if err := object.Delete(ctx); err != nil {
		logger.ErrLog(layers.Firebase, "error deleting object", err)
		return err
	}

	return nil 
}

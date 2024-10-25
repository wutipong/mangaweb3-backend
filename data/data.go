package data

import (
	"context"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"github.com/wutipong/mangaweb3-backend/configuration"
)

var minioClient *minio.Client
var bucket string

func Init(ctx context.Context) error {
	endpoint := configuration.Get().MinIoEndPoint
	accessKey := configuration.Get().MinIoAccessKey
	accessSecret := configuration.Get().MinIoAcessKeySecret
	secure := configuration.Get().MinIoSecure
	bucket = configuration.Get().MinIoBucket

	if c, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: secure,
	}); err != nil {
		log.Fatal().
			Err(err).
			Msg("Unable to create MinIO client.")
		return err
	} else {
		minioClient = c
	}

	if found, err := minioClient.BucketExists(ctx, bucket); !found || err != nil {
		if !found {
			log.Fatal().Msg("MinIO bucket not found.")
			return err
		}

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("MinIO error.")

			return err
		}
	}

	return nil
}

// ListAllItem returns a list all available item.
func ListAllItem(ctx context.Context) (files []string, err error) {
	objectCh := minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			err = object.Err
			return
		}

		ext := strings.ToLower(filepath.Ext(object.Key))

		if ext == ".zip" || ext == ".cbz" {
			files = append(files, object.Key)
		}
	}

	return
}

// ListObject returns a list all child objects.
func ListObject(ctx context.Context, parent string) (files []string, err error) {
	opts := minio.ListObjectsOptions{
		Prefix:    parent,
		Recursive: true,
	}

	// If prefix ends with .zip, add a slash.
	if strings.HasSuffix(parent, ".zip") {
		opts.Prefix = parent + "/"
	}
	opts.Set("x-minio-extract", "true")

	objectCh := minioClient.ListObjects(ctx, bucket, opts)

	for object := range objectCh {
		if object.Err != nil {
			err = object.Err
			return
		}

		ext := strings.ToLower(filepath.Ext(object.Key))

		if ext == ".jpeg" || ext == ".jpg" || ext == ".png" || ext == ".webp" {
			files = append(files, object.Key)
		}
	}

	return
}

// GetObject returns an object reader.
func GetObject(ctx context.Context, name string) (reader io.ReadCloser, err error) {
	opts := minio.GetObjectOptions{}

	opts.Set("x-minio-extract", "true")

	object, err := minioClient.GetObject(ctx, bucket, name, opts)

	if err != nil {
		return
	}

	reader = object

	return
}

func GetLastModifiedTime(ctx context.Context, name string) (lastMod time.Time, err error) {
	objInfo, err := minioClient.StatObject(context.Background(), "mybucket", "myobject", minio.StatObjectOptions{})
	if err != nil {
		return
	}

	lastMod = objInfo.LastModified

	return
}

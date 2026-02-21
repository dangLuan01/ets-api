package s3

import (
	"context"
	"io"
)

type S3Service interface {
	UploadFile(ctx context.Context, bucket, key string, file io.Reader) error
	GetFile(ctx context.Context, bucket, key string) ([]byte, error)
}
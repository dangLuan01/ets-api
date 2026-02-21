package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dangLuan01/ets-api/internal/utils"
)

type s3Service struct {
	clientS3 *s3.Client
}

func NewS3Service(clientS3 *s3.Client) S3Service {
	return &s3Service{
		clientS3: clientS3,
	}
}

func (ss *s3Service) UploadFile(ctx context.Context, bucket, key string, file io.Reader) error {
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		// ContentType: aws.String("image/jpeg"),
	}

	_, err := ss.clientS3.PutObject(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
func (ss *s3Service) GetFile(ctx context.Context, bucket, key string) ([]byte, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	}

	out, err := ss.clientS3.GetObject(ctx, input)
	if err != nil {
		return nil, utils.WrapError(string(utils.ErrCodeInternal), "Cannot get file", err)
	}

	defer out.Body.Close()

	return io.ReadAll(out.Body)
}


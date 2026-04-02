package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/dangLuan01/ets-api/internal/utils"
)

type S3Config struct {
	BucketName      string
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
}

func NewS3Client() *s3.Client {
	cfgS3 := S3Config{
		BucketName: utils.GetEnv("CLOUDFLARE_BUCKETNAME", "mybucket"),
		AccountID: utils.GetEnv("CLOUDFLARE_ACCOUNTID", ""),
		AccessKeyID: utils.GetEnv("CLOUDFLARE_ACCESSEYID", ""),
		AccessKeySecret: utils.GetEnv("CLOUDFLARE_ACCESSKEYSECRET", ""),
	}

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfgS3.AccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfgS3.AccessKeyID, cfgS3.AccessKeySecret, "")),
		config.WithRegion("apac"),
	)

	if err != nil {
		log.Fatalf("Unable connect to S3:%s", err)
	}

	return s3.NewFromConfig(cfg)
}
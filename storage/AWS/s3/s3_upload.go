package s3

import (
	"bytes"
	"context"

	"github.com/ayush5588/push2Storage/internal/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	client *s3.Client
	bucket string
}

func Client(accountInfo map[string]string) (*s3Client, error) {
	c := credentials.NewStaticCredentialsProvider(accountInfo["accessKey"], accountInfo["secretKeyID"], "")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(c),
		config.WithRegion(accountInfo["region"]))

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &s3Client{client, accountInfo["bucket"]}, nil

}

// UploadToStorage read the file from the given path and uploads to s3
func (s3Client *s3Client) UploadToStorage(filePath string) error {
	// Read the file from the given path into a new file.
	// Upload the file to the bucket.

	fileBuffer, fileName, err := util.PrepareFile(filePath)
	if err != nil {
		return err
	}

	uploadClient := manager.NewUploader(s3Client.client)

	_, err = uploadClient.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Client.bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileBuffer),
	})
	if err != nil {
		return err
	}

	return nil

}

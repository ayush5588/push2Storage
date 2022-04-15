package s3

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


type s3Client struct {
	client *s3.Client
	bucket	string
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

func createFile(desiredFileName string) (*os.File, error) {
	var file *os.File
	var err error

	if desiredFileName != "" {
		file, err = os.Create(desiredFileName)
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.Create("file " + time.Now().String())
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

// UploadToStorage read the file from the given path and uploads to s3
func (s3Client *s3Client) UploadToStorage(desiredFileName string, filePath string) (error) {
	// Read the file from the given path into a new file. 
	// Upload the file to the bucket. 
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	
	file, err := createFile(desiredFileName)
	if err != nil {
		return err
	}

	defer file.Close()


	_, err = file.WriteString(string(fileContent))
	if err != nil {
		return err
	}



	uploadClient := manager.NewUploader(s3Client.client)

	result, err := uploadClient.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Client.bucket),
		Key:    aws.String(file.Name()),
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
	
}
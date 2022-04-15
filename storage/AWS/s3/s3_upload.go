package s3

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Client struct {
	s3.Client
}

func Client(accountInfo map[string]string) (*s3Client) {
	c := credentials.NewStaticCredentialsProvider(accountInfo["accessKey"], accountInfo["secretKeyID"], "")
	


}


func (s3Client *s3Client) UploadToStorage(fileName string, fileToUpload []byte) (error) {

}
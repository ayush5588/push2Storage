package upload

import (
	"errors"
	"fmt"

	"github.com/ayush5588/push2Storage/storage/AWS/s3"
)



type uploadResponse struct {
	statuscode int64
	message    string
	err        error
}

func prepareResp(statuscode int64, message string, err error) (uploadResponse) {
	return uploadResponse{
		statuscode,
		message,
		err,
	}
}

// Upload func calls the appropriate method for given storage type in order to upload the given file
func Upload(storageType string, accountInfo map[string]string, desiredFileName string, filePath string) uploadResponse {
	fmt.Println("Check")
	if storageType == "" {
		return prepareResp(400,"", errors.New("no storage type specified"))
	}

	if filePath == "" {
		return prepareResp(400,"", errors.New("no file to upload"))
	}

	switch storageType {
	case "aws":
		s3Client, err := s3.Client(accountInfo)
		if err != nil {
			fmt.Println(err)
			return prepareResp(500,"", err)
		}
		err = s3Client.UploadToStorage(desiredFileName, filePath)
		if err != nil {
			fmt.Println(err)
			return prepareResp(500,"", err)
		}
	}

	return prepareResp(200, "file uploaded successfully", nil)
}
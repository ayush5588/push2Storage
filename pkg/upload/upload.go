package upload

import (
	"errors"
	"fmt"

	"github.com/ayush5588/push2Storage/internal/util"

	"github.com/ayush5588/push2Storage/storage/AWS/s3"
)


var (
	ErrEmptyStorageType error = errors.New("storage field cannot be empty")
	ErrEmptyFilePath error = errors.New("file path cannot be empty")
)




// Upload func calls the appropriate method for given storage type in order to upload the given file
func Upload(storageType string, accountInfo map[string]string, filePath string) (util.UploadResponse) {
	if storageType == "" {
		return util.PrepareResp(400,"", ErrEmptyStorageType, "storage field cannot be empty")
	}

	if filePath == "" {
		return util.PrepareResp(400,"", ErrEmptyFilePath, "file path cannot be empty")
	}

	switch storageType {
	case "aws":
		s3Client, err := s3.Client(accountInfo)
		if err != nil {
			fmt.Println(err)
			return util.PrepareResp(500,"", err, "error in generating a S3 client")
		}
		err = s3Client.UploadToStorage(filePath)
		if err != nil {
			fmt.Println(err)
			return util.PrepareResp(500,"", err, "error in uploading file to S3")
		}
	
	default:
		return util.PrepareResp(404, "implementation for storage type " + storageType + " yet to be done", nil, "")
	}

	return util.PrepareResp(200, "file uploaded successfully", nil, "")
}
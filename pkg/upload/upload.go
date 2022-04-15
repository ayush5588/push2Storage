package upload

import (
	"errors"

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
func Upload(storageType string, accountInfo map[string]string, fileName string, fileToUpload []byte) uploadResponse {
	if storageType == "" {
		return prepareResp(400,"", errors.New("No storage type specified"))
	}

	if len(fileToUpload) == 0 {
		return prepareResp(400,"", errors.New("No file to upload"))
	}

	switch storageType {
	case "aws":
		s3Client := s3.Client(accountInfo)
		err := s3Client.UploadToStorage(fileName, fileToUpload)
		if err != nil {
			return prepareResp(500,"", err)
		}
	}

	return prepareResp(200, "file uploaded successfully", nil)
}
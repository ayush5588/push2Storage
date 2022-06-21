package util

// UploadResponse represents the response metadata body
type UploadResponse struct {
	Statuscode int64
	Message    string
	Error      error `json:"-"` // Not marshalled when sending back json response
	ErrorMsg   string
}

// PrepareResp prepares the response with the metadata
func PrepareResp(statuscode int64, message string, err error, errMsg string) UploadResponse {
	return UploadResponse{
		statuscode,
		message,
		err,
		errMsg,
	}
}
 

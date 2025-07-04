package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayush5588/push2Storage/pkg/upload"
)

func uploadToS3(w http.ResponseWriter, r *http.Request) {
	creds := map[string]string{
		"accessKey":   "<your-access-key>",
		"secretKeyID": "<your-secretKeyID>",
		"bucket":      "<bucket-name>",
		"region":      "<region>",
	}

	result := upload.Upload("aws", creds, `full-path-to-file`)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func main() {
	http.HandleFunc("/", uploadToS3)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayush5588/push2Storage/pkg/upload"
)


func uploadToS3(w http.ResponseWriter, r *http.Request) {
	mp := make(map[string]string)
	mp["accessKey"] = "<your-access-key>"
	mp["secretKeyID"] = "<your-secretKeyID>"
	mp["bucket"] = "<bucket-name>"
	mp["region"] = "<region>"


	result := upload.Upload("aws", mp, "<Name of the file that you want your uploaded file on S3 to be>", `full-path-to-file`)
	
	// eg: result := upload.Upload("aws", mp, "requestdata.txt", `/home/user/test.txt`)
	// The above command will take the file from the given path (/home/user/test.txt), create a new file with name requestdata.txt, copy the data
	// from the original file to this file and upload the requestdata.txt file to S3. 
	
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(result) 


} 

func main() {
	http.HandleFunc("/",uploadToS3)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080",nil)

}

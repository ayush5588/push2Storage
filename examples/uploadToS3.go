package main

import (
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


	result := upload.Upload("aws", mp, "random.txt", `full-path-to-file`)
	fmt.Println(result)
	fmt.Fprintf(w, "Success!!")

} 

func main() {
	http.HandleFunc("/",uploadToS3)
	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080",nil)

}
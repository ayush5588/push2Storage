package util

import (
	"errors"
	"io/ioutil"
	"os"
)




// extractFileName extracts the name of the file along with extension from the given file path
func extractFileName(filePath string) string {

	lenFilePath := len([]rune(filePath))

	var fileName string

	for i := lenFilePath - 1; i >= 0; i-- {
		if filePath[i] == '\\' || filePath[i] == '/' {
			break
		} else {
			fileName = string(filePath[i]) + fileName
		}
	}

	return fileName

}

// prepareFileBuffer prepares buffer with original file content for aws upload client 
func prepareFileBuffer(file *os.File) ([]byte, error) {

	file, err := os.Open(file.Name())
    if err != nil {
        return nil, err
    }


    defer file.Close()


    fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// get the size of the file
    size := fileInfo.Size()
    fileBuffer := make([]byte, size)

	// read the file contents into a buffer
    readFileSize, err := file.Read(fileBuffer)

	if readFileSize != int(size) {
		return nil, errors.New("file not read completely")
	}
	if err != nil {
		return nil, err
	}

	return fileBuffer, nil
}



func PrepareFile(filePath string) ([]byte, string, error) {

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, "", err
	}

	var file *os.File


	fileName := extractFileName(filePath)
	if err != nil {
		return nil, "", err
	}

	file, err = os.Create(fileName)
	if err != nil {
		return nil, "", err
	}

	_, err = file.WriteString(string(fileContent))
	if err != nil {
		return nil, "", err
	}

	fileBuffer, err := prepareFileBuffer(file) 
	if err != nil {
		return nil, "", err
	}

	return fileBuffer, fileName, nil
	
}
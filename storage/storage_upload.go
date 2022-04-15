package storage


type StorageUpload interface {
	UploadToStorage(string, string) (error) 
}
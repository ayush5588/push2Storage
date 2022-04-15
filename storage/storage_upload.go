package storage


type StorageUpload interface {
	UploadToStorage(string, []byte) (error) 
}
package storage

type IStorage interface {
	// Upload Local File and upload it to specific location in cloud service
	//
	// This function return the uploaded file link
	AddCloudFileAndGetURL(localFilePath string) (urlLink string, err error)
}

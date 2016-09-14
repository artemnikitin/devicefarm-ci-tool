package tools

import (
	"log"
	"strings"
)

// UploadFile used to upload file by S3 pre-signed URL
func UploadFile(path, url string) int {
	log.Println("Uploading file from path:", path)
	file, info := prepareFile(path)
	resp := sendRequest(url, &file, info)
	return getStatusOfUpload(resp)
}

// GetFileName returns file name from path string
func GetFileName(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}
	pos := strings.LastIndex(path, "/")
	return string(path[pos+1:])
}

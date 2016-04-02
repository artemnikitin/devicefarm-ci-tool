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

// StringEndsWith check that string ends with specified substring
func StringEndsWith(original, substring string) bool {
	if len(substring) > len(original) {
		return false
	}
	str := string(original[len(original)-len(substring):])
	return str == substring
}

// GetFilename returns file name from path string
func GetFilename(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}
	pos := strings.LastIndex(path, "/")
	return string(path[pos+1:])
}

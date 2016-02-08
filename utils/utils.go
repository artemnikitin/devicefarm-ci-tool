package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// UploadFile used to upload file by S3 pre-signed URL
func UploadFile(path, url string) int {
	log.Println("Uploading app ...")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to get file for upload because of: ", err.Error())
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get info about file because of: ", err.Error())
	}
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, file)
	if err != nil {
		log.Fatal("Failed to create HTTP request because of: ", err.Error())
	}
	request.Header.Add("Content-Type", "application/octet-stream")
	request.ContentLength = info.Size()
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Failed to upload file by S3 link because of: ", err.Error())
	}
	defer resp.Body.Close()
	var result int
	result = resp.StatusCode
	log.Println("Response code:", result)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to get body of response because of: ", err.Error())
	}
	log.Println("Response body:", string(body))
	return result
}

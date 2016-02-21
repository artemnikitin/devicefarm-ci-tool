package tools

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func sendRequest(url string, file *os.File, info os.FileInfo) *http.Response {
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
	return resp
}

func prepareFile(path string) (*os.File, os.FileInfo) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to get file for upload because of: ", err.Error())
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get info about file because of: ", err.Error())
	}
	return file, info
}

func getStatusOfUpload(response *http.Response) int {
	result := response.StatusCode
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to get body of response because of: ", err.Error())
	}
	log.Println("Response code:", result)
	log.Println("Response body:", string(body))
	defer response.Body.Close()
	return result
}

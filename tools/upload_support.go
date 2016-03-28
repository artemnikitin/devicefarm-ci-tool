package tools

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
)

func sendRequest(url string, file *os.File, info os.FileInfo) *http.Response {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, file)
	errors.Validate(err, "Failed to create HTTP request")
	request.Header.Add("Content-Type", "application/octet-stream")
	request.ContentLength = info.Size()
	resp, err := client.Do(request)
	errors.Validate(err, "Failed to upload file by S3 link")
	return resp
}

func prepareFile(path string) (*os.File, os.FileInfo) {
	file, err := os.Open(path)
	errors.Validate(err, "Failed to get file for upload")
	info, err := file.Stat()
	errors.Validate(err, "Failed to get info about file")
	return file, info
}

func getStatusOfUpload(response *http.Response) int {
	result := response.StatusCode
	body, err := ioutil.ReadAll(response.Body)
	errors.Validate(err, "Failed to get body of response")
	log.Println("Response code:", result)
	log.Println("Response body:", string(body))
	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(ioutil.Discard, response.Body)
		response.Body.Close()
	}()
	return result
}

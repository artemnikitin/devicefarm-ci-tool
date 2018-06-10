package tools

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
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
	return path[pos+1:]
}

// GenerateReportURL generate URL to test report from ARN
func GenerateReportURL(arn string) string {
	URL := "https://us-west-2.console.aws.amazon.com/devicefarm/home?region=us-west-2#/projects/%s/runs/%s"
	index := strings.Index(arn, ":run:")
	if index == -1 {
		log.Println("Can't generate test report URL from ARN:", arn)
		return ""
	}
	str := arn[index+5:]
	index = strings.Index(str, "/")
	project := str[:index]
	run := str[index+1:]
	return fmt.Sprintf(URL, project, run)
}

// Random generates random integer in given range
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

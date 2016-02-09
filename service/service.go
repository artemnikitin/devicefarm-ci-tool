package service

import (
	"log"
	"strings"
	"time"

	"github.com/artemnikitin/devicefarm-ci-tool/utils"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/aws"
)

// GetAccountArn returns project ARN by project name
func GetAccountArn(client *devicefarm.DeviceFarm, project string) string {
	var arn string
	params := &devicefarm.ListProjectsInput{}
	resp, err := client.ListProjects(params)
	if err != nil {
		log.Fatal("Failed to get list of projects for account because of: ", err.Error())
	}
	for _, entry := range resp.Projects {
		if *entry.Name == project {
			arn = *entry.Arn
		}
	}
	log.Println("Project ARN:", arn)
	return arn
}

// CreateUpload creates pre-signed S3 URL for upload
func CreateUpload(client *devicefarm.DeviceFarm, arn, appPath string) (string, string) {
	var appType string
	if utils.StringEndsWith(appPath, ".apk") {
		appType = "ANDROID_APP"
	} else {
		appType = "IOS_APP"
	}
	params := &devicefarm.CreateUploadInput{
		Name:        aws.String(getFilename(appPath)),
		ProjectArn:  aws.String(arn),
		Type:        aws.String(appType),
		ContentType: aws.String("application/octet-stream"),
	}
	resp, err := client.CreateUpload(params)
	if err != nil {
		log.Fatal("Failed to upload an app because of: ", err.Error())
	}
	log.Println("Upload ARN:", *resp.Upload.Arn)
	log.Println("Upload URL:", *resp.Upload.Url)
	return *resp.Upload.Arn, *resp.Upload.Url
}

// GetDevicePoolArn returns device pool ARN by device pool name
func GetDevicePoolArn(client *devicefarm.DeviceFarm, projectArn, devicePool string) string {
	var arn string
	params := &devicefarm.ListDevicePoolsInput{
		Arn: aws.String(projectArn),
	}
	resp, err := client.ListDevicePools(params)
	if err != nil {
		log.Fatal("Failed to get list of device pools because of: ", err.Error())
	}
	for _, pool := range resp.DevicePools {
		if *pool.Name == devicePool {
			arn = *pool.Arn
		}
	}
	log.Println("Device pool ARN:", arn)
	return arn
}

// Run creates test run
func Run(client *devicefarm.DeviceFarm, devicePoolArn, projectArn, appArn string) (string, string) {
	log.Println("Starting job ...")
	params := &devicefarm.ScheduleRunInput{
		DevicePoolArn: aws.String(devicePoolArn),
		ProjectArn:    aws.String(projectArn),
		Test: &devicefarm.ScheduleRunTest{
			Type: aws.String("BUILTIN_FUZZ"),
		},
		AppArn: aws.String(appArn),
	}
	resp, err := client.ScheduleRun(params)
	if err != nil {
		log.Fatal("Failed to run tests because of: ", err.Error())
	}
	return *resp.Run.Arn, *resp.Run.Status
}

// WaitForAppProcessed wait while app be in status "SUCCEEDED"
func WaitForAppProcessed(client *devicefarm.DeviceFarm, arn string) {
	var counter int
	status := GetUploadStatus(client, arn)
	for status != "SUCCEEDED" {
		counter++
		time.Sleep(time.Second * 2)
		status = GetUploadStatus(client, arn)
		if status == "FAILED" {
			log.Fatal("Something went wrong with processing app for tests. Quit.")
		}
		if counter == 90 {
			log.Fatal("App is still unprocessed. Quit.")
		}
	}
}

// GetUploadStatus returns status of upload file
func GetUploadStatus(client *devicefarm.DeviceFarm, arn string) string {
	params := &devicefarm.GetUploadInput{
		Arn: aws.String(arn),
	}
	resp, err := client.GetUpload(params)
	if err != nil {
		log.Fatal("Failed to get status of upload because of: ", err.Error())
	}
	log.Println("Status of upload:", *resp.Upload.Status)
	return *resp.Upload.Status
}

func getFilename(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}
	pos := strings.LastIndex(path, "/")
	return string(path[pos+1:])
}

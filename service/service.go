package service

import (
	"log"
	"time"

	"github.com/artemnikitin/devicefarm-ci-tool/config"
	"github.com/artemnikitin/devicefarm-ci-tool/errors"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

// GetProjectArn returns project ARN by project name
func GetProjectArn(client *devicefarm.DeviceFarm, project string) string {
	var arn string
	params := &devicefarm.ListProjectsInput{}
	resp, err := client.ListProjects(params)
	errors.Validate(err, "Failed to get list of projects for account")
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
	if tools.StringEndsWith(appPath, ".apk") {
		appType = "ANDROID_APP"
	} else {
		appType = "IOS_APP"
	}
	return internalCreateUpload(client, arn, appPath, appType)
}

// CreateUploadWithType creates upload with specific type
func CreateUploadWithType(client *devicefarm.DeviceFarm, arn, appPath, uploadType string) (string, string) {
	return internalCreateUpload(client, arn, appPath, uploadType)
}

func internalCreateUpload(client *devicefarm.DeviceFarm, arn, appPath, appType string) (string, string) {
	params := &devicefarm.CreateUploadInput{
		Name:        aws.String(tools.GetFilename(appPath)),
		ProjectArn:  aws.String(arn),
		Type:        aws.String(appType),
		ContentType: aws.String("application/octet-stream"),
	}
	resp, err := client.CreateUpload(params)
	errors.Validate(err, "Failed to upload an app")
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
	errors.Validate(err, "Failed to get list of device pools")
	for _, pool := range resp.DevicePools {
		if *pool.Name == devicePool {
			arn = *pool.Arn
		}
	}
	log.Println("Device pool ARN:", arn)
	return arn
}

// RunWithConfig will schedule run with setup from JSON config
func RunWithConfig(client *devicefarm.DeviceFarm, devicePoolArn, projectArn, appArn string, conf config.RunConfig) (string, string) {
	params := createScheduleRunInput(client, conf, projectArn)
	params.DevicePoolArn = aws.String(devicePoolArn)
	params.AppArn = aws.String(appArn)
	log.Println("Starting job ...")
	resp, err := client.ScheduleRun(params)
	errors.Validate(err, "Failed to run tests")
	log.Println("Run ARN:", *resp.Run.Arn)
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
	errors.Validate(err, "Failed to get status of upload")
	log.Println("Status of upload:", *resp.Upload.Status)
	return *resp.Upload.Status
}

// WaitForRunEnds for run to finish and returns it's result
func WaitForRunEnds(client *devicefarm.DeviceFarm, arn string) {
	status, result := GetStatusOfRun(client, arn)
	for status != "COMPLETED" {
		log.Println("Waiting for run to finish...")
		time.Sleep(time.Second * 5)
		status, result = GetStatusOfRun(client, arn)
	}
	if result == "PASSED" {
		log.Println("Run successfully finished!")
	} else {
		log.Fatal("There was a problem with this run. Status :", result)
	}
}

// GetStatusOfRun returns status and result of run specified by ARN
func GetStatusOfRun(client *devicefarm.DeviceFarm, arn string) (string, string) {
	params := &devicefarm.GetRunInput{
		Arn: aws.String(arn),
	}
	resp, err := client.GetRun(params)
	errors.Validate(err, "Can't get status of run")
	return *resp.Run.Status, *resp.Run.Result
}

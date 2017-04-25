package service

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
)

// GetProjectArn returns project ARN by project name
func GetProjectArn(client devicefarmiface.DeviceFarmAPI, project string) string {
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
func CreateUpload(client devicefarmiface.DeviceFarmAPI, arn, appPath string) (string, string) {
	var appType string
	if strings.HasSuffix(appPath, ".apk") {
		appType = "ANDROID_APP"
	} else {
		appType = "IOS_APP"
	}
	return internalCreateUpload(client, arn, appPath, appType)
}

// CreateUploadWithType creates upload with specific type
func CreateUploadWithType(client devicefarmiface.DeviceFarmAPI, arn, appPath, uploadType string) (string, string) {
	return internalCreateUpload(client, arn, appPath, uploadType)
}

func internalCreateUpload(client devicefarmiface.DeviceFarmAPI, arn, appPath, appType string) (string, string) {
	params := &devicefarm.CreateUploadInput{
		Name:        aws.String(tools.GetFileName(appPath)),
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
func GetDevicePoolArn(client devicefarmiface.DeviceFarmAPI, projectArn, devicePool string) string {
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

// RunWithConfig will schedule run with setup from JSON model
func RunWithConfig(p *model.RunParameters) (string, string) {
	params := createScheduleRunInput(p)
	params.DevicePoolArn = aws.String(p.DeviceArn)
	params.AppArn = aws.String(p.AppArn)
	log.Println("Starting job ...")
	resp, err := p.Client.ScheduleRun(params)
	errors.Validate(err, "Failed to run tests")
	log.Println("Run ARN:", *resp.Run.Arn)
	return *resp.Run.Arn, *resp.Run.Status
}

// WaitForAppProcessed wait while app be in status "SUCCEEDED"
func WaitForAppProcessed(client devicefarmiface.DeviceFarmAPI, arn string, timeout int) {
	var counter int
	limit := 300 / timeout
	status := GetUploadStatus(client, arn)
	for status != "SUCCEEDED" {
		counter++
		time.Sleep(time.Duration(timeout) * time.Second)
		status = GetUploadStatus(client, arn)
		if status == "FAILED" {
			log.Fatal("Something went wrong with processing app for tests. Quit.")
		}
		if counter == limit {
			log.Fatal("App is still unprocessed. Quit.")
		}
	}
}

// GetUploadStatus returns status of upload file
func GetUploadStatus(client devicefarmiface.DeviceFarmAPI, arn string) string {
	params := &devicefarm.GetUploadInput{
		Arn: aws.String(arn),
	}
	resp, err := client.GetUpload(params)
	errors.Validate(err, "Failed to get status of upload")
	log.Println("Status of upload:", *resp.Upload.Status)
	return *resp.Upload.Status
}

// WaitForRunEnds for run to finish and returns it's result
func WaitForRunEnds(client devicefarmiface.DeviceFarmAPI, arn string, checkEvery int) string {
	status, result := GetStatusOfRun(client, arn)
	for status != "COMPLETED" {
		log.Println("Waiting for run to finish...")
		time.Sleep(time.Duration(checkEvery) * time.Second)
		status, result = GetStatusOfRun(client, arn)
	}
	log.Println("Run finished!")
	return result
}

// GetStatusOfRun returns status and result of run specified by ARN
func GetStatusOfRun(client devicefarmiface.DeviceFarmAPI, arn string) (string, string) {
	params := &devicefarm.GetRunInput{
		Arn: aws.String(arn),
	}
	resp, err := client.GetRun(params)
	errors.Validate(err, "Can't get status of run")
	return *resp.Run.Status, *resp.Run.Result
}

// GetListOfFailedTests returns list of failed test for a specified run with additional info
func GetListOfFailedTests(client devicefarmiface.DeviceFarmAPI, arn string) []*model.FailedTest {
	var wg sync.WaitGroup
	var m sync.Mutex
	var result []*model.FailedTest

	jobs := getListOfJobsForRun(client, arn)

	wg.Add(len(jobs))
	for i := 0; i < len(jobs); i++ {
		go func(i int) {
			device := *jobs[i].Name
			os := *jobs[i].Device.Platform + " " + *jobs[i].Device.Os

			suites := getListOfSuitesForJob(client, *jobs[i].Arn)
			suitesArn := getListOfTestArnFromSuite(suites)
			tests := getListOfFailedTestsFromSuite(client, suitesArn, device, os)
			tempResult := populateResult(tests, client)

			m.Lock()
			result = append(result, tempResult...)
			m.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	return result
}

func getListOfJobsForRun(client devicefarmiface.DeviceFarmAPI, arn string) []*devicefarm.Job {
	params := &devicefarm.ListJobsInput{
		Arn: aws.String(arn),
	}
	resp, err := client.ListJobs(params)
	errors.Validate(err, "Can't get list of jobs for run")
	return resp.Jobs
}

func getListOfSuitesForJob(client devicefarmiface.DeviceFarmAPI, arn string) []*devicefarm.Suite {
	params := &devicefarm.ListSuitesInput{
		Arn: aws.String(arn),
	}
	resp, err := client.ListSuites(params)
	errors.Validate(err, "Can't get list of suites for job")
	return resp.Suites
}

func getListOfTestForSuite(client devicefarmiface.DeviceFarmAPI, arn string) []*devicefarm.Test {
	params := &devicefarm.ListTestsInput{
		Arn: aws.String(arn),
	}
	resp, err := client.ListTests(params)
	errors.Validate(err, "Can't get list of tests for suite")
	return resp.Tests
}

func getArtifactsForTest(client devicefarmiface.DeviceFarmAPI, arn string) []*devicefarm.Artifact {
	params := &devicefarm.ListArtifactsInput{
		Arn:  aws.String(arn),
		Type: aws.String(devicefarm.ArtifactCategoryFile),
	}
	resp, err := client.ListArtifacts(params)
	errors.Validate(err, "Can't get list of artifacts for a test")
	return resp.Artifacts
}

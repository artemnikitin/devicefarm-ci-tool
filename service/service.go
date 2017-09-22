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

// DeviceFarmRun represents parameters for utility runtime
type DeviceFarmRun struct {
	Client     devicefarmiface.DeviceFarmAPI
	Config     *model.RunConfig
	Project    string
	ProjectArn string
	DeviceArn  string
	AppArn     string
}

// GetProjectArn returns project ARN by project name
func (p *DeviceFarmRun) GetProjectArn() string {
	var arn string
	params := &devicefarm.ListProjectsInput{}
	resp, err := p.Client.ListProjects(params)
	errors.Validate(err, "Failed to get list of projects for account")
	for _, entry := range resp.Projects {
		if *entry.Name == p.Project {
			arn = *entry.Arn
		}
	}
	log.Println("Project ARN:", arn)
	p.ProjectArn = arn
	return arn
}

// CreateUpload creates pre-signed S3 URL for upload
func (p *DeviceFarmRun) CreateUpload(appPath string) (string, string) {
	var appType string
	if strings.HasSuffix(appPath, ".apk") {
		appType = devicefarm.UploadTypeAndroidApp
	} else {
		appType = devicefarm.UploadTypeIosApp
	}
	return internalCreateUpload(p.Client, p.ProjectArn, appPath, appType)
}

// CreateUploadWithType creates upload with specific type
func (p *DeviceFarmRun) CreateUploadWithType(arn, appPath, uploadType string) (string, string) {
	return internalCreateUpload(p.Client, arn, appPath, uploadType)
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
func (p *DeviceFarmRun) GetDevicePoolArn(devicePool string) string {
	if p.Config.DevicePoolArn != "" {
		return p.Config.DevicePoolArn
	}
	var arn string
	params := &devicefarm.ListDevicePoolsInput{
		Arn: aws.String(p.ProjectArn),
	}
	resp, err := p.Client.ListDevicePools(params)
	errors.Validate(err, "Failed to get list of device pools")
	for _, pool := range resp.DevicePools {
		if *pool.Name == devicePool {
			arn = *pool.Arn
		}
	}
	log.Println("Device pool ARN:", arn)
	p.DeviceArn = arn
	p.Config.DevicePoolArn = arn
	return arn
}

// RunWithConfig will schedule run with setup from JSON model
func (p *DeviceFarmRun) RunWithConfig() (string, string) {
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
func (p *DeviceFarmRun) WaitForAppProcessed(arn string, timeout int) {
	var counter int
	limit := 300 / timeout
	status := p.GetUploadStatus(arn)
	for status != devicefarm.UploadStatusSucceeded {
		counter++
		time.Sleep(time.Duration(timeout) * time.Second)
		status = p.GetUploadStatus(arn)
		if status == devicefarm.UploadStatusFailed {
			log.Fatal("Something went wrong with processing app for tests. Quit.")
		}
		if counter == limit {
			log.Fatal("App is still unprocessed. Quit.")
		}
	}
}

// GetUploadStatus returns status of upload file
func (p *DeviceFarmRun) GetUploadStatus(arn string) string {
	params := &devicefarm.GetUploadInput{
		Arn: aws.String(arn),
	}
	resp, err := p.Client.GetUpload(params)
	errors.Validate(err, "Failed to get status of upload")
	log.Println("Status of upload:", *resp.Upload.Status)
	return *resp.Upload.Status
}

// WaitForRunEnds for run to finish and returns it's result
func (p *DeviceFarmRun) WaitForRunEnds(arn string, checkEvery int) string {
	status, result := p.GetStatusOfRun(arn)
	for status != devicefarm.ExecutionStatusCompleted {
		log.Println("Waiting for run to finish...")
		time.Sleep(time.Duration(checkEvery) * time.Second)
		status, result = p.GetStatusOfRun(arn)
	}
	log.Println("Run finished!")
	return result
}

// GetStatusOfRun returns status and result of run specified by ARN
func (p *DeviceFarmRun) GetStatusOfRun(arn string) (string, string) {
	params := &devicefarm.GetRunInput{
		Arn: aws.String(arn),
	}
	resp, err := p.Client.GetRun(params)
	errors.Validate(err, "Can't get status of run")
	return *resp.Run.Status, *resp.Run.Result
}

// GetListOfFailedTests returns list of failed test for a specified run with additional info
func (p *DeviceFarmRun) GetListOfFailedTests(arn string) []*model.FailedTest {
	var wg sync.WaitGroup
	var m sync.Mutex
	var failedTests []*model.FailedTest

	jobs := getListOfJobsForRun(p.Client, arn)

	wg.Add(len(jobs))
	for i := 0; i < len(jobs); i++ {
		go func(i int) {
			device := *jobs[i].Name
			os := *jobs[i].Device.Platform + " " + *jobs[i].Device.Os

			suites := getListOfSuitesForJob(p.Client, *jobs[i].Arn)
			suitesArn := getListOfTestArnFromSuite(suites)
			tests := getListOfFailedTestsFromSuite(p.Client, suitesArn, device, os)
			tempResult := populateResult(tests, p.Client)

			m.Lock()
			failedTests = append(failedTests, tempResult...)
			m.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	return failedTests
}

// IsTestRunPassIgnoringUnavailableDevices checks if there was a situation then test runs on some devices passes,
// but some devices weren't available. In this case it returns 'true', otherwise 'false'
func (p *DeviceFarmRun) IsTestRunPassIgnoringUnavailableDevices(arn string) bool {
	result := false

	jobs := getListOfJobsForRun(p.Client, arn)

	for _, v := range jobs {
		if *v.Result != devicefarm.ExecutionResultPassed && *v.Result != devicefarm.ExecutionResultErrored {
			return false
		}

		suites := getListOfSuitesForJob(p.Client, *v.Arn)

		for _, k := range suites {
			tests := getListOfTestForSuite(p.Client, *k.Arn)
			for i := range tests {
				if *tests[i].Name != "Setup Test" && *tests[i].Name != "Teardown Test" {
					if *tests[i].Result == devicefarm.ExecutionResultPassed {
						result = true
					}
				}
			}
		}
	}

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

package service

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/artemnikitin/devicefarm-ci-tool/config"
	"github.com/artemnikitin/devicefarm-ci-tool/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"sync"
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
	return internalCreateUpload(client, arn, appPath, appType)
}

// CreateUploadWithType creates upload with specific type
func CreateUploadWithType(client *devicefarm.DeviceFarm, arn, appPath, uploadType string) (string, string) {
	return internalCreateUpload(client, arn, appPath, uploadType)
}

func internalCreateUpload(client *devicefarm.DeviceFarm, arn, appPath, appType string) (string, string) {
	params := &devicefarm.CreateUploadInput{
		Name:        aws.String(utils.GetFilename(appPath)),
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

// Run creates test run with default parameters
func Run(client *devicefarm.DeviceFarm, devicePoolArn, projectArn, appArn string) (string, string) {
	params := &devicefarm.ScheduleRunInput{
		DevicePoolArn: aws.String(devicePoolArn),
		ProjectArn:    aws.String(projectArn),
		Test: &devicefarm.ScheduleRunTest{
			Type: aws.String("BUILTIN_FUZZ"),
		},
		AppArn: aws.String(appArn),
	}
	return runWith(client, params)
}

// RunWithConfig will schedule run with setup from JSON config
func RunWithConfig(client *devicefarm.DeviceFarm, devicePoolArn, projectArn, appArn string, conf config.RunConfig) (string, string) {
	params := createScheduleRunInput(client, conf, projectArn)
	params.DevicePoolArn = aws.String(devicePoolArn)
	params.AppArn = aws.String(appArn)
	params.ProjectArn = aws.String(projectArn)
	return runWith(client, params)
}

func runWith(client *devicefarm.DeviceFarm, input *devicefarm.ScheduleRunInput) (string, string) {
	log.Println("Starting job ...")
	resp, err := client.ScheduleRun(input)
	if err != nil {
		log.Fatal("Failed to run tests because of: ", err.Error())
	}
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
	if err != nil {
		log.Fatal("Failed to get status of upload because of: ", err.Error())
	}
	log.Println("Status of upload:", *resp.Upload.Status)
	return *resp.Upload.Status
}

func createScheduleRunInput(client *devicefarm.DeviceFarm, conf config.RunConfig, projectArn string) *devicefarm.ScheduleRunInput {
	var wg sync.WaitGroup
	result := &devicefarm.ScheduleRunInput{
		Test: &devicefarm.ScheduleRunTest{},
		Configuration: &devicefarm.ScheduleRunConfiguration{
			Radios: &devicefarm.Radios{
				Bluetooth: aws.Bool(true),
				Gps:       aws.Bool(true),
				Nfc:       aws.Bool(true),
				Wifi:      aws.Bool(true),
			},
		},
	}
	if conf.RunName != "" {
		result.Name = aws.String(conf.RunName)
	}
	if conf.Test.Type != "" {
		result.Test.Type = aws.String(conf.Test.Type)
	}
	if conf.Test.TestPackageArn == "" && conf.Test.TestPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare tests for uploading...")
			t := config.GetUploadTypeForTest(conf.Test.Type)
			arn, url := CreateUploadWithType(client, projectArn, conf.Test.TestPackagePath, t)
			httpResponse := utils.UploadFile(conf.Test.TestPackagePath, url)
			if httpResponse != 200 {
				log.Fatal("Can't upload test app")
			}
			WaitForAppProcessed(client, arn)
			result.Test.TestPackageArn = aws.String(arn)
			wg.Done()
		}()
	}
	if conf.Test.TestPackageArn != "" {
		result.Test.TestPackageArn = aws.String(conf.Test.TestPackageArn)
	}
	if conf.Test.Filter != "" {
		result.Test.Filter = aws.String(conf.Test.Filter)
	}
	params := conf.Test.Parameters
	if len(params) > 0 {
		temp := make(map[string]*string)
		for k, v := range params {
			temp[k] = aws.String(v)
		}
		result.Test.Parameters = temp
	}
	if conf.AdditionalData.BillingMethod != "" {
		result.Configuration.BillingMethod = aws.String(conf.AdditionalData.BillingMethod)
	}
	if conf.AdditionalData.Locale != "" {
		result.Configuration.Locale = aws.String(conf.AdditionalData.Locale)
	}
	if conf.AdditionalData.NetworkProfileArn != "" {
		result.Configuration.NetworkProfileArn = aws.String(conf.AdditionalData.NetworkProfileArn)
	}
	if conf.AdditionalData.Location.Latitude != 0 && conf.AdditionalData.Location.Longitude != 0 {
		result.Configuration.Location.Latitude = aws.Float64(conf.AdditionalData.Location.Latitude)
		result.Configuration.Location.Longitude = aws.Float64(conf.AdditionalData.Location.Longitude)
	}
	if len(conf.AdditionalData.AuxiliaryApps) != 0 {
		array := conf.AdditionalData.AuxiliaryApps
		result.Configuration.AuxiliaryApps = aws.StringSlice(array)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Bluetooth) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Bluetooth)
		result.Configuration.Radios.Bluetooth = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Gps) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Gps)
		result.Configuration.Radios.Gps = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Nfc) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Nfc)
		result.Configuration.Radios.Nfc = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Wifi) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Wifi)
		result.Configuration.Radios.Wifi = aws.Bool(b)
	}
	if conf.AdditionalData.ExtraDataPackageArn != "" {
		result.Configuration.ExtraDataPackageArn = aws.String(conf.AdditionalData.ExtraDataPackageArn)
	}
	if conf.AdditionalData.ExtraDataPackageArn == "" && conf.AdditionalData.ExtraDataPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare extra data for uploading...")
			arn, url := CreateUploadWithType(client, projectArn, conf.AdditionalData.ExtraDataPackagePath, "EXTERNAL_DATA")
			httpResponse := utils.UploadFile(conf.AdditionalData.ExtraDataPackagePath, url)
			if httpResponse != 200 {
				log.Fatal("Can't upload test app")
			}
			WaitForAppProcessed(client, arn)
			result.Configuration.ExtraDataPackageArn = aws.String(arn)
			wg.Done()
		}()
	}
	wg.Wait()
	return result
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
	if err != nil {
		log.Fatal("Can't get status of run because of:", err)
	}
	return *resp.Run.Status, *resp.Run.Result
}

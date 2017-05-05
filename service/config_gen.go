package service

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

func createScheduleRunInput(p *DeviceFarmRun) *devicefarm.ScheduleRunInput {
	var wg sync.WaitGroup
	result := &devicefarm.ScheduleRunInput{
		ProjectArn: aws.String(p.ProjectArn),
		Test:       &devicefarm.ScheduleRunTest{},
		Configuration: &devicefarm.ScheduleRunConfiguration{
			Radios: &devicefarm.Radios{
				Bluetooth: aws.Bool(true),
				Gps:       aws.Bool(true),
				Nfc:       aws.Bool(true),
				Wifi:      aws.Bool(true),
			},
			Location: &devicefarm.Location{
				Latitude:  aws.Float64(47.6204),
				Longitude: aws.Float64(-122.3491),
			},
		},
		ExecutionConfiguration: &devicefarm.ExecutionConfiguration{
			JobTimeoutMinutes:  aws.Int64(60),
			AccountsCleanup:    aws.Bool(false),
			AppPackagesCleanup: aws.Bool(false),
		},
	}

	result.Name = aws.String(p.Config.RunName)
	result.Configuration.AuxiliaryApps = aws.StringSlice(p.Config.AdditionalData.AuxiliaryApps)
	if p.Config.AdditionalData.BillingMethod != "" {
		result.Configuration.BillingMethod = aws.String(p.Config.AdditionalData.BillingMethod)
	}
	result.Configuration.Locale = aws.String(p.Config.AdditionalData.Locale)
	if p.Config.AdditionalData.Location.Latitude != 0 {
		result.Configuration.Location.Latitude = aws.Float64(p.Config.AdditionalData.Location.Latitude)
	}
	if p.Config.AdditionalData.Location.Longitude != 0 {
		result.Configuration.Location.Longitude = aws.Float64(p.Config.AdditionalData.Location.Longitude)
	}
	if p.Config.AdditionalData.NetworkProfileArn != "" {
		result.Configuration.NetworkProfileArn = aws.String(p.Config.AdditionalData.NetworkProfileArn)
	}
	result.Configuration.Radios.Bluetooth = aws.Bool(stringToBool(p.Config.AdditionalData.Radios.Bluetooth))
	result.Configuration.Radios.Gps = aws.Bool(stringToBool(p.Config.AdditionalData.Radios.Gps))
	result.Configuration.Radios.Nfc = aws.Bool(stringToBool(p.Config.AdditionalData.Radios.Nfc))
	result.Configuration.Radios.Wifi = aws.Bool(stringToBool(p.Config.AdditionalData.Radios.Wifi))
	result.Test.Filter = aws.String(p.Config.Test.Filter)
	result.Test.Parameters = aws.StringMap(p.Config.Test.Parameters)
	if p.Config.Test.Type != "" {
		result.Test.Type = aws.String(p.Config.Test.Type)
	} else {
		result.Test.Type = aws.String(devicefarm.TestTypeBuiltinFuzz)
	}
	if p.Config.Test.TestPackageArn != "" {
		result.Test.TestPackageArn = aws.String(p.Config.Test.TestPackageArn)
	} else {
		uploadTestPackage(p, result, &wg)
	}
	if p.Config.AdditionalData.ExtraDataPackageArn != "" {
		result.Configuration.ExtraDataPackageArn = aws.String(p.Config.AdditionalData.ExtraDataPackageArn)
	} else {
		uploadExtraData(p, result, &wg)
	}
	if p.Config.ExecutionConfiguration.JobTimeoutMinutes > 0 {
		result.ExecutionConfiguration.JobTimeoutMinutes = aws.Int64(int64(p.Config.ExecutionConfiguration.JobTimeoutMinutes))
	}
	if p.Config.ExecutionConfiguration.AccountsCleanup {
		result.ExecutionConfiguration.AccountsCleanup = aws.Bool(true)
	}
	if p.Config.ExecutionConfiguration.AppPackagesCleanup {
		result.ExecutionConfiguration.AppPackagesCleanup = aws.Bool(true)
	}
	wg.Wait()
	return result
}

func stringToBool(str string) bool {
	b, err := strconv.ParseBool(strings.ToLower(str))
	if err != nil {
		return true
	}
	return b
}

func uploadExtraData(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	if p.Config.AdditionalData.ExtraDataPackageArn == "" && p.Config.AdditionalData.ExtraDataPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare extra data for uploading...")
			arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.AdditionalData.ExtraDataPackagePath, "EXTERNAL_DATA")
			httpResponse := tools.UploadFile(p.Config.AdditionalData.ExtraDataPackagePath, url)
			if httpResponse != 200 {
				log.Fatal("Can't upload test app")
			}
			p.WaitForAppProcessed(arn, 5)
			result.Configuration.ExtraDataPackageArn = aws.String(arn)
			wg.Done()
		}()
	}
}

func uploadTestPackage(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	if p.Config.Test.TestPackageArn == "" && p.Config.Test.TestPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare tests for uploading...")
			t := model.GetUploadTypeForTest(p.Config.Test.Type)
			arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.Test.TestPackagePath, t)
			httpResponse := tools.UploadFile(p.Config.Test.TestPackagePath, url)
			if httpResponse != 200 {
				log.Fatal("Can't upload test app")
			}
			p.WaitForAppProcessed(arn, 5)
			result.Test.TestPackageArn = aws.String(arn)
			wg.Done()
		}()
	}
}

package service

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/artemnikitin/devicefarm-ci-tool/config"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

func createScheduleRunInput(client *devicefarm.DeviceFarm, conf *config.RunConfig, projectArn string) *devicefarm.ScheduleRunInput {
	var wg sync.WaitGroup
	result := &devicefarm.ScheduleRunInput{
		ProjectArn: aws.String(projectArn),
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
	}

	result.Name = aws.String(conf.RunName)
	result.Configuration.AuxiliaryApps = aws.StringSlice(conf.AdditionalData.AuxiliaryApps)
	if conf.AdditionalData.BillingMethod != "" {
		result.Configuration.BillingMethod = aws.String(conf.AdditionalData.BillingMethod)
	}
	result.Configuration.Locale = aws.String(conf.AdditionalData.Locale)
	if conf.AdditionalData.Location.Latitude != 0 {
		result.Configuration.Location.Latitude = aws.Float64(conf.AdditionalData.Location.Latitude)
	}
	if conf.AdditionalData.Location.Longitude != 0 {
		result.Configuration.Location.Longitude = aws.Float64(conf.AdditionalData.Location.Longitude)
	}
	if conf.AdditionalData.NetworkProfileArn != "" {
		result.Configuration.NetworkProfileArn = aws.String(conf.AdditionalData.NetworkProfileArn)
	}
	result.Configuration.Radios.Bluetooth = aws.Bool(stringToBool(conf.AdditionalData.Radios.Bluetooth))
	result.Configuration.Radios.Gps = aws.Bool(stringToBool(conf.AdditionalData.Radios.Gps))
	result.Configuration.Radios.Nfc = aws.Bool(stringToBool(conf.AdditionalData.Radios.Nfc))
	result.Configuration.Radios.Wifi = aws.Bool(stringToBool(conf.AdditionalData.Radios.Wifi))
	result.Test.Filter = aws.String(conf.Test.Filter)
	result.Test.Parameters = aws.StringMap(conf.Test.Parameters)
	if conf.Test.Type != "" {
		result.Test.Type = aws.String(conf.Test.Type)
	} else {
		result.Test.Type = aws.String("BUILTIN_FUZZ")
	}

	if conf.Test.TestPackageArn == "" && conf.Test.TestPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare tests for uploading...")
			t := config.GetUploadTypeForTest(conf.Test.Type)
			arn, url := CreateUploadWithType(client, projectArn, conf.Test.TestPackagePath, t)
			httpResponse := tools.UploadFile(conf.Test.TestPackagePath, url)
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

	if conf.AdditionalData.ExtraDataPackageArn == "" && conf.AdditionalData.ExtraDataPackagePath != "" {
		wg.Add(1)
		go func() {
			log.Println("Prepare extra data for uploading...")
			arn, url := CreateUploadWithType(client, projectArn, conf.AdditionalData.ExtraDataPackagePath, "EXTERNAL_DATA")
			httpResponse := tools.UploadFile(conf.AdditionalData.ExtraDataPackagePath, url)
			if httpResponse != 200 {
				log.Fatal("Can't upload test app")
			}
			WaitForAppProcessed(client, arn)
			result.Configuration.ExtraDataPackageArn = aws.String(arn)
			wg.Done()
		}()
	}
	if conf.AdditionalData.ExtraDataPackageArn != "" {
		result.Configuration.ExtraDataPackageArn = aws.String(conf.AdditionalData.ExtraDataPackageArn)
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

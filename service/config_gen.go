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

func createScheduleRunInput(client *devicefarm.DeviceFarm, conf config.RunConfig, projectArn string) *devicefarm.ScheduleRunInput {
	var wg sync.WaitGroup
	result := &devicefarm.ScheduleRunInput{
		ProjectArn: aws.String(projectArn),
		Test: &devicefarm.ScheduleRunTest{},
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
	if conf.RunName != "" {
		result.Name = aws.String(conf.RunName)
	}

	processTestBlock(conf, result)
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

	processConfigurationBlock(conf, result)
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
	wg.Wait()
	return result
}

func processTestBlock(conf config.RunConfig, sri *devicefarm.ScheduleRunInput) {
	if conf.Test.Type != "" {
		sri.Test.Type = aws.String(conf.Test.Type)
	}
	if conf.Test.TestPackageArn != "" {
		sri.Test.TestPackageArn = aws.String(conf.Test.TestPackageArn)
	}
	if conf.Test.Filter != "" {
		sri.Test.Filter = aws.String(conf.Test.Filter)
	}
	params := conf.Test.Parameters
	if len(params) > 0 {
		temp := make(map[string]*string)
		for k, v := range params {
			temp[k] = aws.String(v)
		}
		sri.Test.Parameters = temp
	}
}

func processConfigurationBlock(conf config.RunConfig, sri *devicefarm.ScheduleRunInput) {
	if conf.AdditionalData.BillingMethod != "" {
		sri.Configuration.BillingMethod = aws.String(conf.AdditionalData.BillingMethod)
	}
	if conf.AdditionalData.Locale != "" {
		sri.Configuration.Locale = aws.String(conf.AdditionalData.Locale)
	}
	if conf.AdditionalData.NetworkProfileArn != "" {
		sri.Configuration.NetworkProfileArn = aws.String(conf.AdditionalData.NetworkProfileArn)
	}
	if conf.AdditionalData.Location.Latitude != 0 && conf.AdditionalData.Location.Longitude != 0 {
		sri.Configuration.Location.Latitude = aws.Float64(conf.AdditionalData.Location.Latitude)
		sri.Configuration.Location.Longitude = aws.Float64(conf.AdditionalData.Location.Longitude)
	}
	if len(conf.AdditionalData.AuxiliaryApps) != 0 {
		array := conf.AdditionalData.AuxiliaryApps
		sri.Configuration.AuxiliaryApps = aws.StringSlice(array)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Bluetooth) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Bluetooth)
		sri.Configuration.Radios.Bluetooth = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Gps) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Gps)
		sri.Configuration.Radios.Gps = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Nfc) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Nfc)
		sri.Configuration.Radios.Nfc = aws.Bool(b)
	}
	if strings.ToLower(conf.AdditionalData.Radios.Wifi) == "false" {
		b, _ := strconv.ParseBool(conf.AdditionalData.Radios.Wifi)
		sri.Configuration.Radios.Wifi = aws.Bool(b)
	}
	if conf.AdditionalData.ExtraDataPackageArn != "" {
		sri.Configuration.ExtraDataPackageArn = aws.String(conf.AdditionalData.ExtraDataPackageArn)
	}
}

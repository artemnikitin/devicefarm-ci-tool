package service

import (
	"log"
	"sync"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/fatih/structs"
)

func createScheduleRunInput(p *DeviceFarmRun) *devicefarm.ScheduleRunInput {
	var wg sync.WaitGroup

	result := &devicefarm.ScheduleRunInput{
		ProjectArn: aws.String(p.ProjectArn),
	}

	config := structs.New(p.Config)

	v, ok := config.FieldOk("Name")
	if ok && !v.IsZero() {
		result.Name = aws.String(v.Value().(string))
	}

	v, ok = config.FieldOk("ExecutionConfiguration")
	if ok && !v.IsZero() {
		result.ExecutionConfiguration = p.Config.ExecutionConfiguration
	}

	v, ok = config.FieldOk("Test")
	if ok && !v.IsZero() {
		result.Test = p.Config.Test
		arn, arnOK := v.FieldOk("TestPackageArn")
		path, pathOK := config.FieldOk("TestPackagePath")
		if ((arnOK && arn.IsZero()) || !arnOK) && pathOK && !path.IsZero() {
			uploadTestPackage(p, result, &wg)
		}
	} else {
		result.Test = &devicefarm.ScheduleRunTest{
			Type: aws.String(devicefarm.TestTypeBuiltinFuzz),
		}
	}

	v, ok = config.FieldOk("Configuration")
	if ok && !v.IsZero() {
		result.Configuration = p.Config.Configuration
		arn, arnOK := v.FieldOk("ExtraDataPackageArn")
		path, pathOK := config.FieldOk("ExtraDataPackagePath")
		if ((arnOK && arn.IsZero()) || !arnOK) && pathOK && !path.IsZero() {
			uploadExtraData(p, result, &wg)
		}
	}

	wg.Wait()
	return result
}

func uploadExtraData(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		log.Println("Prepare extra data for uploading...")
		arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.ExtraDataPackagePath, devicefarm.UploadTypeExternalData)
		httpResponse := tools.UploadFile(p.Config.ExtraDataPackagePath, url)
		if httpResponse != 200 {
			log.Fatal("Can't upload test app")
		}
		p.WaitForAppProcessed(arn, 5)
		result.Configuration.ExtraDataPackageArn = aws.String(arn)
		wg.Done()
	}()
}

func uploadTestPackage(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		log.Println("Prepare tests for uploading...")
		t := model.GetUploadTypeForTest(*p.Config.Test.Type)
		arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.TestPackagePath, t)
		httpResponse := tools.UploadFile(p.Config.TestPackagePath, url)
		if httpResponse != 200 {
			log.Fatal("Can't upload test app")
		}
		p.WaitForAppProcessed(arn, 5)
		result.Test.TestPackageArn = aws.String(arn)
		wg.Done()
	}()
}

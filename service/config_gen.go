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
		s := structs.New(result)
		v, ok = s.Field("Test").FieldOk("Type")
		if !ok || v.IsZero() {
			result.Test.Type = aws.String(devicefarm.TestTypeBuiltinFuzz)
		}
	} else {
		result.Test = &devicefarm.ScheduleRunTest{
			Type: aws.String(devicefarm.TestTypeBuiltinFuzz),
		}
	}

	s := structs.New(result)
	v, ok = s.FieldOk("Test")
	if ok {
		arn, arnOK := v.FieldOk("TestPackageArn")
		path, pathOK := config.FieldOk("TestPackagePath")
		if ((arnOK && arn.IsZero()) || !arnOK) && pathOK && !path.IsZero() {
			uploadTestPackage(p, result, &wg)
		}
	}

	v, ok = config.FieldOk("Configuration")
	if ok && !v.IsZero() {
		result.Configuration = p.Config.Configuration
	}

	s = structs.New(result)
	path, pathOK := config.FieldOk("ExtraDataPackagePath")
	if pathOK && !path.IsZero() {
		v, ok = s.FieldOk("Configuration")
		if ok && !v.IsZero() {
			arn, arnOK := v.FieldOk("ExtraDataPackageArn")
			if (arnOK && arn.IsZero()) || !arnOK {
				uploadExtraData(p, result, &wg)
			}
		} else {
			result.Configuration = &devicefarm.ScheduleRunConfiguration{}
			uploadExtraData(p, result, &wg)
		}
	}

	if len(p.Config.AuxiliaryAppsPath) > 0 {
		auxAppARN := make([]string, 0)
		for _, v := range p.Config.AuxiliaryAppsPath {
			auxAppARN = append(auxAppARN, uploadAuxiliaryApps(v, p))
		}
		f, ok := structs.New(result).FieldOk("Configuration")
		if !ok || (ok && f.IsZero()) {
			result.Configuration = &devicefarm.ScheduleRunConfiguration{}
		}
		result.Configuration.AuxiliaryApps = aws.StringSlice(auxAppARN)
	}
	specPath, specPathOK := config.FieldOk("TestSpecPath")
	if specPathOK && !specPath.IsZero() {
		v, ok = s.FieldOk("Configuration")
		if ok && !v.IsZero() {
			arn, arnOK := v.FieldOk("TestSpecPath")
			if (arnOK && arn.IsZero()) || !arnOK {
				uploadTestSpec(p, result, &wg)
			}
		} else {
			result.Configuration = &devicefarm.ScheduleRunConfiguration{}
			uploadTestSpec(p, result, &wg)
		}
	}

	wg.Wait()
	return result
}

func uploadExtraData(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		log.Println("Preparing extra data for upload...")
		arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.ExtraDataPackagePath, devicefarm.UploadTypeExternalData)
		httpResponse := tools.UploadFile(p.Config.ExtraDataPackagePath, url)
		if httpResponse != 200 {
			log.Fatal("Can't upload extra data")
		}
		p.WaitForAppProcessed(arn, 5)
		result.Configuration.ExtraDataPackageArn = aws.String(arn)
		wg.Done()
	}()
}

func uploadTestPackage(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		log.Println("Preparing tests for upload...")
		t := model.GetUploadTypeForTest(*result.Test.Type)
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

func uploadAuxiliaryApps(app string, p *DeviceFarmRun) string {
	log.Println("Uploading auxiliary app...")
	arn, url := p.CreateUpload(app)
	httpResponse := tools.UploadFile(app, url)
	if httpResponse != 200 {
		log.Fatal("Can't upload auxiliary app")
	}
	p.WaitForAppProcessed(arn, 5)
	return arn
}

func uploadTestSpec(p *DeviceFarmRun, result *devicefarm.ScheduleRunInput, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		log.Println("Uploading testspec...")
		t := model.GetUploadTypeForTestSpec(*result.Test.Type)
		arn, url := p.CreateUploadWithType(p.ProjectArn, p.Config.TestSpecPath, t)
		httpResponse := tools.UploadFile(p.Config.TestSpecPath, url)
		if httpResponse != 200 {
			log.Fatal("Can't upload test app")
		}
		p.WaitForAppProcessed(arn, 5)
		result.Test.TestSpecArn = aws.String(arn)
		wg.Done()
	}()
}

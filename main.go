package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/service"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
)

var (
	project                  = flag.String("project", "", "Device Farm project name")
	runName                  = flag.String("run", "", "Name of test run")
	appPath                  = flag.String("app", "", "Path to an app")
	testPath                 = flag.String("test", "", "Path to test app")
	devicePool               = flag.String("devices", "Top Devices", "Specify list of devices for tests")
	randomDevicePool         = flag.String("randomDevices", "", "List of device pools for random choice")
	configJSON               = flag.String("config", "", "Path to JSON config")
	wait                     = flag.Bool("wait", false, "Wait for run end")
	checkEvery               = flag.Int("checkEvery", 5, "Specified time slice for checking status of run")
	useRandomDevicePool      = flag.Bool("useRandomDevicePool", false, "If enabled will select one of available device pools")
	ignoreUnavailableDevices = flag.Bool("ignoreUnavailableDevices", false, "Consider test run where one of devices failed as green")
	testType                 = flag.String("testType", "", "Type of tests to run")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	*project = "iot-sdk"
	*appPath = "HereUiKitApp-debug.apk"
	*useRandomDevicePool = true

	if *project == "" || *appPath == "" {
		fmt.Println("Please specify correct parameters!")
		fmt.Println("You should specify:")
		fmt.Println("-project, name of a project in AWS Device Farm")
		fmt.Println("-app, path to an app you want to run")
		os.Exit(1)
	}

	failedTests, success := runJob(getAWSClient(), getConfig())
	if !success {
		if len(failedTests) == 0 {
			log.Println(fmt.Sprintf("Test job fails, it looks like AWS/infrastructure/setup issue, check the report."))
		} else {
			log.Println(fmt.Sprintf("There are %d test fails, check it out!\n", len(failedTests)))
			for i := 0; i < len(failedTests); i++ {
				fmt.Println(failedTests[i].ToString())
			}
		}
		os.Exit(1)
	}
}

func runJob(client devicefarmiface.DeviceFarmAPI, config *model.RunConfig) ([]*model.FailedTest, bool) {
	svc := &service.DeviceFarmRun{
		Client:  client,
		Config:  config,
		Project: *project,
	}

	if svc.GetProjectArn() == "" {
		log.Fatal("Application finished, because it can't retrieve project ARN")
	}

	if *useRandomDevicePool {
		svc.SelectDevicePool(svc.ProjectArn)
	}
	svc.GetDevicePoolArn(svc.Config.DevicePoolName)

	appArn, url := svc.CreateUpload(*appPath)
	code := tools.UploadFile(*appPath, url)
	if code != 200 {
		log.Fatal("Can't upload an app to Device Farm")
	}

	svc.AppArn = appArn
	svc.WaitForAppProcessed(svc.AppArn, *checkEvery)

	runArn, status := svc.RunWithConfig()
	statusCheck(status)

	pass := true
	var failedTests []*model.FailedTest

	if *wait {
		result := svc.WaitForRunEnds(runArn, *checkEvery)
		if result != devicefarm.ExecutionResultPassed {
			failed := svc.GetListOfFailedTests(runArn)
			pass = makeBuildFailed(result, failed, *ignoreUnavailableDevices)
		}
	}

	if *ignoreUnavailableDevices && *wait {
		pass = svc.IsTestRunPassIgnoringUnavailableDevices(runArn)
	}

	printReportURL(runArn)
	return failedTests, pass
}

func getConfig() *model.RunConfig {
	configFile := &model.RunConfig{
		Test: &devicefarm.ScheduleRunTest{},
	}
	if *configJSON != "" {
		bytes, err := ioutil.ReadFile(*configJSON)
		errors.Validate(err, "Can't read model file")
		configFile = model.Transform(bytes)
	}
	if configFile.DevicePoolArn == "" && configFile.DevicePoolName == "" {
		if *randomDevicePool != "" {
			pools := strings.Split(*randomDevicePool, ",")
			i := tools.Random(0, len(pools)-1)
			configFile.DevicePoolName = pools[i]
		} else {
			configFile.DevicePoolName = *devicePool
		}
	}
	if *runName != "" {
		configFile.Name = *runName
	}
	if *testPath != "" {
		configFile.TestPackagePath = *testPath
	}
	if *testType != "" {
		configFile.Test.Type = aws.String(*testType)
	}
	return configFile
}

func getAWSClient() devicefarmiface.DeviceFarmAPI {
	config := aws.NewConfig()
	config.WithCredentials(credentials.NewEnvCredentials())
	config.WithRegion("us-west-2")
	ses, err := session.NewSession(config)
	errors.Validate(err, "Can't create an AWS session")
	return devicefarm.New(ses)
}

func statusCheck(status string) {
	if status == devicefarm.ExecutionStatusScheduling {
		log.Println("Job is scheduled!")
	} else {
		log.Println("Status =", status)
		log.Fatal("Failed to start a job ...")
	}
}

func makeBuildFailed(result string, failed []*model.FailedTest, option bool) bool {
	testResult := false
	if result == devicefarm.ExecutionResultErrored && len(failed) == 0 && option {
		testResult = true
	}
	return testResult
}

func printReportURL(arn string) {
	log.Println("Test report URL:", tools.GenerateReportURL(arn))
}

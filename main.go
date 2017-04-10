package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/artemnikitin/aws-config"
	"github.com/artemnikitin/devicefarm-ci-tool/errors"
	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/artemnikitin/devicefarm-ci-tool/service"
	"github.com/artemnikitin/devicefarm-ci-tool/tools"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	project    = flag.String("project", "", "Device Farm project name")
	runName    = flag.String("run", "", "Name of test run")
	appPath    = flag.String("app", "", "Path to an app")
	testPath   = flag.String("test", "", "Path to test app")
	devicePool = flag.String("devices", "Top Devices", "Specify list of devices for tests")
	configJSON = flag.String("config", "", "Path to JSON config")
	wait       = flag.Bool("wait", false, "Wait for run end")
	checkEvery = flag.Int("checkEvery", 5, "Specified time slice for checking status of run")
)

func main() {
	flag.Parse()

	if *project == "" || *appPath == "" {
		fmt.Println("Please specify correct parameters!")
		fmt.Println("You should specify:")
		fmt.Println("-project, name of a project in AWS Device Farm")
		fmt.Println("-app, path to an app you want to run")
		os.Exit(1)
	}

	runJob(getAWSClient(), getConfig())
}

func runJob(client *devicefarm.DeviceFarm, config *model.RunConfig) {
	p := &model.RunParameters{
		Client:  client,
		Config:  config,
		Project: *project,
	}
	p.ProjectArn = service.GetProjectArn(p.Client, p.Project)
	if p.ProjectArn == "" {
		log.Fatal("Application finished, because it can't retrieve project ARN")
	}
	p.DeviceArn = service.GetDevicePoolArn(p.Client, p.ProjectArn, *devicePool)
	appArn, url := service.CreateUpload(p.Client, p.ProjectArn, *appPath)
	code := tools.UploadFile(*appPath, url)
	if code != 200 {
		log.Fatal("Can't upload an app to Device Farm")
	}
	p.AppArn = appArn
	service.WaitForAppProcessed(p.Client, p.AppArn)
	runArn, status := service.RunWithConfig(p)
	statusCheck(status)
	if *wait {
		result := service.WaitForRunEnds(p.Client, runArn, *checkEvery)
		printReportURL(runArn)
		if result != "PASSED" {
			fails := service.GetListOfFailedTests(client, runArn)
			log.Printf("There are %d test fails, check it out!\n", len(fails))
			for i := 0; i < len(fails); i++ {
				fmt.Println(fails[i].ToString())
			}
			os.Exit(1)
		}
	}
}

func getConfig() *model.RunConfig {
	configFile := &model.RunConfig{}
	if *configJSON != "" {
		bytes, err := ioutil.ReadFile(*configJSON)
		errors.Validate(err, "Can't read model file")
		*configFile = model.Transform(bytes)
	}
	if *runName != "" {
		configFile.RunName = *runName
	}
	if *testPath != "" {
		configFile.Test.TestPackagePath = *testPath
	}
	return configFile
}

func getAWSClient() *devicefarm.DeviceFarm {
	config := awsconfig.New()
	if *config.Region != "us-west-2" {
		config.WithRegion("us-west-2")
	}
	session := session.New(config)
	return devicefarm.New(session)
}

func statusCheck(status string) {
	if status == "SCHEDULING" {
		log.Println("Job is scheduled!")
	} else {
		log.Println("Status =", status)
		log.Fatal("Failed to start a job ...")
	}
}

func printReportURL(arn string) {
	log.Println("Test report URL:", tools.GenerateReportURL(arn))
}

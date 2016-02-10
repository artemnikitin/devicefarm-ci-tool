package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/artemnikitin/aws-config"
	"github.com/artemnikitin/devicefarm-ci-tool/config"
	"github.com/artemnikitin/devicefarm-ci-tool/service"
	"github.com/artemnikitin/devicefarm-ci-tool/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	project    = flag.String("project", "", "Device Farm project name")
	appPath    = flag.String("app", "", "Path to an app")
	devicePool = flag.String("devices", "Top Devices", "Specify list of devices for tests")
	configJSON = flag.String("config", "", "Path to JSON config")
	wait       = flag.Bool("wait", false, "Wait for run end")
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

	var configFile config.RunConfig
	if *configJSON != "" {
		bytes, err := ioutil.ReadFile(*configJSON)
		if err != nil {
			log.Fatal("Can't read config file because of:", err)
		}
		configFile = config.Transform(bytes)
	}

	config := awsconfig.New()
	if *config.Region != "us-west-2" {
		config.WithRegion("us-west-2")
	}
	session := session.New(config)
	client := devicefarm.New(session)

	projectArn := service.GetAccountArn(client, *project)
	if projectArn != "" {
		deviceArn := service.GetDevicePoolArn(client, projectArn, *devicePool)
		appArn, url := service.CreateUpload(client, projectArn, *appPath)
		code := utils.UploadFile(*appPath, url)
		if code != 200 {
			log.Fatal("Can't upload an app to Device Farm")
		}
		service.WaitForAppProcessed(client, appArn)
		var status string
		var runArn string
		if *configJSON != "" {
			runArn, status = service.RunWithConfig(client, deviceArn, projectArn, appArn, configFile)
		} else {
			runArn, status = service.Run(client, deviceArn, projectArn, appArn)
		}
		if status == "SCHEDULING" {
			log.Println("Job is started!")
		} else {
			log.Println("Status =", status)
			log.Fatal("Failed to start a job ...")
		}
		if *wait {
			service.WaitForRunEnds(client, runArn)
		}
	}
}

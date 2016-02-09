package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/artemnikitin/aws-config"
	"github.com/artemnikitin/devicefarm-ci-tool/service"
	"github.com/artemnikitin/devicefarm-ci-tool/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	project    = flag.String("project", "", "Device Farm project name")
	appPath    = flag.String("app", "", "Path to an app")
	devicePool = flag.String("devices", "Top Devices", "Specify list of devices for tests")
	//configJSON = flag.String("config", "", "Path to JSON config")
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
		if code == 200 {
			service.WaitForAppProcessed(client, appArn)
			_, status := service.Run(client, deviceArn, projectArn, appArn)
			if status == "SCHEDULING" {
				log.Println("Job is started!")
			}
		}
	}
}
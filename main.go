package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	project    = flag.String("project", "", "Device Farm project name")
	appPath    = flag.String("app", "", "Path to an app")
	devicePool = flag.String("devices", "Top Devices", "Specify list of devices for tests")
	logging    = flag.Bool("log", false, "Enable logging of HTTP requests to AWS")
	region     = flag.String("region", "us-west-2", "Set AWS region")
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

	session := session.New(createConfig())
	client := devicefarm.New(session)

	projectArn := getAccountArn(client)
	if projectArn != "" {
		deviceArn := getDevicePoolArn(client, projectArn)
		appArn, url := createUpload(client, projectArn)
		code := uploadFile(url)
		if code == 200 {
			waitForAppProcessed(client, appArn)
			_, status := run(client, deviceArn, projectArn, appArn)
			if status == "SCHEDULING" {
				log.Println("Job is started!")
			}
		}
	}
}

func getAccountArn(client *devicefarm.DeviceFarm) string {
	var arn string
	params := &devicefarm.ListProjectsInput{}
	resp, err := client.ListProjects(params)
	if err != nil {
		log.Fatal("Failed to get list of projects for account because of: ", err.Error())
	}
	for _, entry := range resp.Projects {
		if *entry.Name == *project {
			arn = *entry.Arn
		}
	}
	log.Println("Project ARN:", arn)
	return arn
}

func createUpload(client *devicefarm.DeviceFarm, arn string) (string, string) {
	var appType string
	if stringEndsWith(*appPath, ".apk") {
		appType = "ANDROID_APP"
	} else {
		appType = "IOS_APP"
	}
	params := &devicefarm.CreateUploadInput{
		Name:        aws.String(getFilename(*appPath)),
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

func getDevicePoolArn(client *devicefarm.DeviceFarm, projectArn string) string {
	var arn string
	params := &devicefarm.ListDevicePoolsInput{
		Arn: aws.String(projectArn),
	}
	resp, err := client.ListDevicePools(params)
	if err != nil {
		log.Fatal("Failed to get list of device pools because of: ", err.Error())
	}
	for _, pool := range resp.DevicePools {
		if *pool.Name == *devicePool {
			arn = *pool.Arn
		}
	}
	log.Println("Device pool ARN:", arn)
	return arn
}

func run(client *devicefarm.DeviceFarm, devicePoolArn, projectArn, appArn string) (string, string) {
	log.Println("Starting job ...")
	params := &devicefarm.ScheduleRunInput{
		DevicePoolArn: aws.String(devicePoolArn),
		ProjectArn:    aws.String(projectArn),
		Test: &devicefarm.ScheduleRunTest{
			Type: aws.String("BUILTIN_FUZZ"),
		},
		AppArn: aws.String(appArn),
	}
	resp, err := client.ScheduleRun(params)
	if err != nil {
		log.Fatal("Failed to run tests because of: ", err.Error())
	}
	return *resp.Run.Arn, *resp.Run.Status
}

func getUploadStatus(client *devicefarm.DeviceFarm, arn string) string {
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

func waitForAppProcessed(client *devicefarm.DeviceFarm, arn string) {
	var counter int
	status := getUploadStatus(client, arn)
	for status != "SUCCEEDED" {
		counter++
		time.Sleep(time.Second * 2)
		status = getUploadStatus(client, arn)
		if status == "FAILED" {
			log.Fatal("Something went wrong with processing app for tests. Quit.")
		}
		if counter == 90 {
			log.Fatal("App is still unprocessed. Quit.")
		}
	}
}

func uploadFile(url string) int {
	log.Println("Uploading app ...")
	file, err := os.Open(*appPath)
	if err != nil {
		log.Fatal("Failed to get file for upload because of: ", err.Error())
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get info about file because of: ", err.Error())
	}
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, file)
	if err != nil {
		log.Fatal("Failed to create HTTP request because of: ", err.Error())
	}
	request.Header.Add("Content-Type", "application/octet-stream")
	request.ContentLength = info.Size()
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Failed to upload file by S3 link because of: ", err.Error())
	}
	defer resp.Body.Close()
	var result int
	result = resp.StatusCode
	log.Println("Response code:", result)
	if *logging {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to get body of response because of: ", err.Error())
		}
		log.Println("Response body:", string(body))
	}
	return result
}

func stringEndsWith(original, substring string) bool {
	if len(substring) > len(original) {
		return false
	}
	str := string(original[len(original)-len(substring) : len(original)])
	return str == substring
}

func getFilename(path string) string {
	if !strings.Contains(path, "/") {
		return path
	}
	pos := strings.LastIndex(path, "/")
	return string(path[pos+1:])
}

func createConfig() *aws.Config {
	config := aws.NewConfig()
	config.WithCredentials(credentials.NewEnvCredentials())
	config.WithRegion(*region)
	if *logging {
		config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}
	return config
}

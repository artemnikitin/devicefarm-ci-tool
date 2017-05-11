package service

import (
	"os"
	"testing"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	testRun = &DeviceFarmRun{
		Client:  &MockClient{},
		Config:  &model.RunConfig{},
		Project: os.Getenv("AWS_DEVICE_FARM_PROJECT"),
	}
	testRunFailed = &DeviceFarmRun{
		Client: &MockClient{
			Failed: true,
		},
		Config:  &model.RunConfig{},
		Project: os.Getenv("AWS_DEVICE_FARM_PROJECT"),
	}
)

func TestGetProjectArn(t *testing.T) {
	expected := "qwerty"
	arn := testRun.GetProjectArn()
	if arn != expected {
		t.Errorf("Expected: %s, actual: %s", expected, arn)
	}
	if arn != testRun.ProjectArn {
		t.Errorf("Both ARN should be equal: %s and %s", arn, testRun.ProjectArn)
	}
}

func TestGetProjectArnFailed(t *testing.T) {
	expected := "qwerty"
	arn := testRunFailed.GetProjectArn()
	if arn == expected {
		t.Error("ARN should be different!")
	}
	if arn != testRunFailed.ProjectArn {
		t.Errorf("Both ARN should be equal: %s and %s", arn, testRunFailed.ProjectArn)
	}
}

func TestCreateUpload(t *testing.T) {
	arn, url := testRun.CreateUpload("app.apk")
	if arn == "" {
		t.Error("ARN shouldn't be empty")
	}
	if url == "" {
		t.Error("URL shouldn't be empty")
	}
}

func TestCreateUploadFailed(t *testing.T) {
	arn, url := testRunFailed.CreateUpload("")
	if arn != "" {
		t.Error("ARN should be empty")
	}
	if url != "" {
		t.Error("URL should be empty")
	}
}

func TestGetDevicePoolArn(t *testing.T) {
	expected := "test"
	arn := testRun.GetDevicePoolArn("111")
	if arn != expected {
		t.Errorf("Expected: %s, actual: %s", expected, arn)
	}
}

func TestGetDevicePoolArnFailed(t *testing.T) {
	expected := "test2"
	arn := testRun.GetDevicePoolArn("")
	if arn == expected {
		t.Errorf("Expected: %s, actual: %s", expected, arn)
	}
}

func TestRunWithConfig(t *testing.T) {
	arn, status := testRun.RunWithConfig()
	if arn == "" {
		t.Error("ARN shouldn't be empty string")
	}
	if status == "" {
		t.Error("Status shouldn't be empty string")
	}
}

func TestRunWithConfigFailed(t *testing.T) {
	arn, status := testRunFailed.RunWithConfig()
	if arn != "" {
		t.Error("ARN should be empty string")
	}
	if status != "" {
		t.Error("Status should be empty string")
	}
}

func TestGetUploadStatus(t *testing.T) {
	status := testRun.GetUploadStatus("")
	if status == "" {
		t.Error("Status shouldn't be empty")
	}
}

func TestGetUploadStatusFailed(t *testing.T) {
	status := testRunFailed.GetUploadStatus("")
	if status != "" {
		t.Error("Status should be empty")
	}
}

func TestWaitForRunEnds(t *testing.T) {
	result := testRun.WaitForRunEnds("", 1)
	if result != devicefarm.ExecutionResultPassed {
		t.Error("Result should be ok")
	}
}

func TestGetStatusOfRun(t *testing.T) {
	status, result := testRun.GetStatusOfRun("")
	if status == "" {
		t.Error("Status shouldn't be empty")
	}
	if result == "" {
		t.Error("Result shouldn't be empty")
	}
}

func TestGetStatusOfRunFailed(t *testing.T) {
	status, result := testRunFailed.GetStatusOfRun("")
	if status != "" {
		t.Error("Status should be empty")
	}
	if result != "" {
		t.Error("Result should be empty")
	}
}

func TestGetListOfFailedTests(t *testing.T) {
	failed := testRun.GetListOfFailedTests("")
	if len(failed) != 0 {
		t.Error("List of failed tests should be 0!")
	}
}

func TestGetListOfFailedTestsMoreThan0(t *testing.T) {
	failed := testRunFailed.GetListOfFailedTests("")
	if len(failed) != 2 {
		t.Error("List of failed tests should be 2!")
	}
}

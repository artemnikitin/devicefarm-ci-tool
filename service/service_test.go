package service

import (
	"os"
	"testing"

	"github.com/artemnikitin/aws-config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var (
	project    = os.Getenv("AWS_DEVICE_FARM_PROJECT")
	testClient = devicefarm.New(session.New(awsconfig.New().WithRegion("us-west-2")))
)

func TestGetProjectArnForExistedProject(t *testing.T) {
	result := GetProjectArn(testClient, project)
	if len(result) == 0 {
		t.Error("For existed project ARN should be returned")
	}
}

func TestGetProjectArnForUnexistedProject(t *testing.T) {
	result := GetProjectArn(testClient, "bla-bla-bla2234342")
	if len(result) != 0 {
		t.Error("For unexisted project ARN should be blank string")
	}
}

package service

import (
	"fmt"
	"testing"

	"github.com/artemnikitin/aws-config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var testClient = devicefarm.New(session.New(awsconfig.New()))

func TestA(t *testing.T) {
	fmt.Println(testClient.APIVersion)
}

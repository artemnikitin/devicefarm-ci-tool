package model

import (
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

// RunParameters represents parameters for utility runtime
type RunParameters struct {
	Client     *devicefarm.DeviceFarm
	Config     *RunConfig
	Project    string
	ProjectArn string
	DeviceArn  string
	AppArn     string
}

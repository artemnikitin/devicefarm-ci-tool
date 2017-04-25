package model

import "github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"

// RunParameters represents parameters for utility runtime
type RunParameters struct {
	Client     devicefarmiface.DeviceFarmAPI
	Config     *RunConfig
	Project    string
	ProjectArn string
	DeviceArn  string
	AppArn     string
}

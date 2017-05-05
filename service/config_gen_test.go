package service

import (
	"fmt"
	"testing"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var client = devicefarm.New(session.New(aws.NewConfig()))

func TestGenerateScheduleRunInputWithConfigurationBlock(t *testing.T) {
	input := []byte(`{"runName":"name","test":{"type":"string","testPackageArn":"string","filter":"string","parameters":{"key1":"value","key2":"value"}},"additionalData":{"extraDataPackageArn":"string","networkProfileArn":"string","locale":"string","location":{"latitude":1.222,"longitude":1.222},"radios":{"wifi":"","bluetooth":"true","nfc":"true","gps":"false"},"auxiliaryApps":["string1","string2"],"billingMethod":"METERED"}}`)
	deviceFarmConfig := create(input)
	if *deviceFarmConfig.Configuration.BillingMethod != "METERED" {
		t.Error("billing method should be 'METERED'")
	}
	if *deviceFarmConfig.Configuration.ExtraDataPackageArn != "string" {
		t.Error("extraDataPackageARN should be 'string'")
	}
	if *deviceFarmConfig.Configuration.NetworkProfileArn != "string" {
		t.Error("networkProfileARN should be 'string'")
	}
	if *deviceFarmConfig.Configuration.Locale != "string" {
		t.Error("locale should be 'string'")
	}
	if len(deviceFarmConfig.Configuration.AuxiliaryApps) != 2 {
		t.Error("should be 2 aux apps")
	}
	if *deviceFarmConfig.Configuration.Location.Latitude != 1.222 {
		t.Error("lat should be 1.222")
	}
	if *deviceFarmConfig.Configuration.Location.Longitude != 1.222 {
		t.Error("lon should be 1.222")
	}
	if *deviceFarmConfig.Configuration.Radios.Gps {
		t.Error("gps should be false")
	}
	if !*deviceFarmConfig.Configuration.Radios.Wifi {
		t.Error("wifi should be true")
	}
}

func TestGenerateScheduleRunInputFromEmptyConfig(t *testing.T) {
	input := []byte(`{"runName":"name"}`)
	conf := create(input)
	if *conf.Name != "name" {
		t.Error("Name should be equal 'name'")
	}
	if !*conf.Configuration.Radios.Bluetooth {
		t.Error("Bluetooth should be true by default")
	}
	if *conf.Configuration.Location.Latitude != 47.6204 {
		t.Error("Latitude should be 47.6204 by default")
	}
}

func TestGenerateScheduleRunInputWithTestBlock(t *testing.T) {
	input := []byte(`{"test":{"type":"string","testPackageArn":"string","filter":"string","parameters":{"key1":"value","key2":"value"}}}`)
	conf := create(input)
	if *conf.Test.Filter != "string" {
		t.Error("test.filter should be 'string'")
	}
	if *conf.Test.Type != "string" {
		t.Error("test.type should be 'string'")
	}
	if *conf.Test.TestPackageArn != "string" {
		t.Error("test.packageARN should be 'string'")
	}
	if len(conf.Test.Parameters) != 2 {
		t.Error("Size of test.parameters should be 2")
	}
	if *conf.Test.Parameters["key1"] != "value" {
		t.Error("test.parameters should be 'value'")
	}
	if *conf.Test.Parameters["key2"] != "value" {
		t.Error("test.parameters should be 'value'")
	}
}

func TestStringToBool(t *testing.T) {
	cases := map[string]bool{
		"true":             true,
		"false":            false,
		"TRUE":             true,
		"FALSE":            false,
		"TrUe":             true,
		"FaLsE":            false,
		"incorrect string": true,
	}

	for k, v := range cases {
		res := stringToBool(k)
		if res != v {
			t.Errorf("For case: %s, actual: %t, expected: %t", k, res, v)
		}
	}
}

func TestCheckExecutionConfiguration(t *testing.T) {
	input := []byte(`{"runName":"name"}`)
	conf := create(input)
	if *conf.ExecutionConfiguration.JobTimeoutMinutes != 60 {
		t.Error("Job timeout by default should be nil")
	}
	if *conf.ExecutionConfiguration.AppPackagesCleanup {
		t.Error("Cleanup by default should be false")
	}
	if *conf.ExecutionConfiguration.AccountsCleanup {
		t.Error("Cleanup by default should be false")
	}
}

func create(bytes []byte) *devicefarm.ScheduleRunInput {
	cf := model.Transform(bytes)
	p := &DeviceFarmRun{
		Client:  client,
		Config:  &cf,
		Project: "232323",
	}
	deviceFarmConfig := createScheduleRunInput(p)
	fmt.Println(deviceFarmConfig.String())
	return deviceFarmConfig
}

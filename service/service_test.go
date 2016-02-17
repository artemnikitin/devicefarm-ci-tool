package service

import (
	"fmt"
	"testing"

	"github.com/artemnikitin/aws-config"
	"github.com/artemnikitin/devicefarm-ci-tool/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var client = devicefarm.New(session.New(awsconfig.New()))

func TestGenerateScheduleRunInputFullConfig(t *testing.T) {
	input := []byte(`{"runName":"name","test":{"type":"string","testPackageArn":"string","filter":"string","parameters":{"key1":"value","key2":"value"}},"additionalData":{"extraDataPackageArn":"string","networkProfileArn":"string","locale":"string","location":{"latitude":1.222,"longitude":1.222},"radios":{"wifi":"","bluetooth":"true","nfc":"true","gps":"false"},"auxiliaryApps":["string1","string2"],"billingMethod":"METERED"}}`)
	conf := config.Transform(input)
	deviceFarmConfig := createScheduleRunInput(client, conf, "232323")
	fmt.Println(deviceFarmConfig.String())
	if *deviceFarmConfig.Name != "name" {
		t.Error("Name should be equal 'name'")
	}
	if *deviceFarmConfig.Test.Filter != "string" {
		t.Error("test.filter should be 'string'")
	}
	if *deviceFarmConfig.Test.Type != "string" {
		t.Error("test.type should be 'string'")
	}
	if *deviceFarmConfig.Test.TestPackageArn != "string" {
		t.Error("test.packageARN should be 'string'")
	}
	if len(deviceFarmConfig.Test.Parameters) != 2 {
		t.Error("Size of test.parameters should be 2")
	}
	if *deviceFarmConfig.Test.Parameters["key1"] != "value" {
		t.Error("test.parameters should be 'value'")
	}
	if *deviceFarmConfig.Test.Parameters["key2"] != "value" {
		t.Error("test.parameters should be 'value'")
	}
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

package service

import (
	"fmt"
	"testing"

	"github.com/artemnikitin/devicefarm-ci-tool/model"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
	"github.com/fatih/structs"
)

func TestGenerateScheduleRunInputWithConfigurationBlock(t *testing.T) {
	input := []byte(`{"name":"name","test":{"type":"string","testPackageArn":"string","filter":"string","parameters":{"key1":"value","key2":"value"}},"configuration":{"extraDataPackageArn":"string","networkProfileArn":"string","locale":"string","location":{"latitude":1.222,"longitude":1.222},"radios":{"bluetooth":true,"nfc":true,"gps":false},"auxiliaryApps":["string1","string2"],"billingMethod":"METERED"}}`)
	conf := createRun(input, &MockClient{})
	if *conf.Configuration.BillingMethod != "METERED" {
		t.Error("billing method should be 'METERED'")
	}
	if *conf.Configuration.ExtraDataPackageArn != "string" {
		t.Error("extraDataPackageARN should be 'string'")
	}
	if *conf.Configuration.NetworkProfileArn != "string" {
		t.Error("networkProfileARN should be 'string'")
	}
	if *conf.Configuration.Locale != "string" {
		t.Error("locale should be 'string'")
	}
	if len(conf.Configuration.AuxiliaryApps) != 2 {
		t.Error("should be 2 aux apps")
	}
	if *conf.Configuration.Location.Latitude != 1.222 {
		t.Error("lat should be 1.222")
	}
	if *conf.Configuration.Location.Longitude != 1.222 {
		t.Error("lon should be 1.222")
	}
	if *conf.Configuration.Radios.Gps {
		t.Error("gps should be false")
	}
}

func TestGenerateScheduleRunInputFromEmptyConfig(t *testing.T) {
	input := []byte(`{"name":"name"}`)
	conf := createRun(input, &MockClient{})
	if *conf.Name != "name" {
		t.Error("Name should be equal 'name'")
	}
	if *conf.Test.Type != devicefarm.TestTypeBuiltinFuzz {
		t.Error("Test type should be set by default")
	}
	s := structs.New(conf)
	if !s.Field("Configuration").IsZero() {
		t.Error("Block should be missed")
	}
}

func TestGenerateScheduleRunInputWithTestBlock(t *testing.T) {
	input := []byte(`{"test":{"type":"string","testPackageArn":"string","filter":"string","parameters":{"key1":"value","key2":"value"}}}`)
	conf := createRun(input, &MockClient{})
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

func TestCheckExecutionConfigurationNonDefault(t *testing.T) {
	input := []byte(`{"name":"name", "executionConfiguration":{"jobTimeoutMinutes":11,"accountsCleanup":true,"appPackagesCleanup":true}}`)
	conf := createRun(input, &MockClient{})
	if *conf.ExecutionConfiguration.JobTimeoutMinutes != 11 {
		t.Error("Job timeout for initialized value should be initialized value")
	}
	if !*conf.ExecutionConfiguration.AppPackagesCleanup {
		t.Error("Cleanup for initialized value should be initialized value")
	}
	if !*conf.ExecutionConfiguration.AccountsCleanup {
		t.Error("Cleanup for initialized value should be initialized value")
	}
}

func TestUploadTestPackage(t *testing.T) {
	server := CreateFakeServer()
	defer server.Close()

	client := &MockClient{
		UploadTest: true,
		FakeServer: server,
	}

	cases := []struct {
		name        string
		expectedARN string
		jsonString  []byte
	}{
		{
			name:        "If test package ARN exists, then it should be used",
			jsonString:  []byte(`{"test":{"testPackageArn":"qqqq"}}`),
			expectedARN: "qqqq",
		},
		{
			name:        "If test package path exists and no test ARN, then package should be upload and ARN generated",
			jsonString:  []byte(`{"testPackagePath":"test.zzz"}`),
			expectedARN: UploadARN,
		},
		{
			name:        "If both ARN and path presented, then ARN should be used",
			jsonString:  []byte(`{"testPackagePath":"test.zzz", "test":{"testPackageArn":"qqqq"}}`),
			expectedARN: "qqqq",
		},
		{
			name:        "Both ARN and path missed, then no ARN should be presented",
			jsonString:  []byte(`{}`),
			expectedARN: "",
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			conf := createRun(v.jsonString, client)
			fmt.Println(conf.String())
			s := structs.New(conf)
			f, ok := s.Field("Test").FieldOk("TestPackageArn")
			if !ok && f.IsZero() && v.expectedARN != "" {
				t.Fatalf("Test:%s\n TestPackageArn field isn't exist or has default value", v.name)
			}
			if v.expectedARN == "" && ok && !f.IsZero() {
				t.Fatalf("Test:%s\n TestPackageArn shouldn't be presented if it's not expected", v.name)
			}
			if v.expectedARN != "" && *conf.Test.TestPackageArn != v.expectedARN {
				t.Fatalf("Test:%s\n Expected: %s, actual: %s", v.name, v.expectedARN, *conf.Test.TestPackageArn)
			}
		})
	}
}

func TestUploadExtraData(t *testing.T) {
	server := CreateFakeServer()
	defer server.Close()

	client := &MockClient{
		UploadTest: true,
		FakeServer: server,
	}

	cases := []struct {
		name        string
		expectedARN string
		jsonString  []byte
	}{
		{
			name:        "If ARN exists, then it should be used",
			jsonString:  []byte(`{"configuration":{"extraDataPackageArn":"qqqq"}}`),
			expectedARN: "qqqq",
		},
		{
			name:        "If path exist and no ARN, then package should be upload and ARN generated",
			jsonString:  []byte(`{"extraDataPackagePath":"test.zzz"}`),
			expectedARN: UploadARN,
		},
		{
			name:        "If both ARN and path presented, then ARN should be used",
			jsonString:  []byte(`{"extraDataPackagePath":"test.zzz", "configuration":{"extraDataPackageArn":"qqqq"}}`),
			expectedARN: "qqqq",
		},
		{
			name:        "If path exist and no ARN, then package should be upload and ARN generated",
			jsonString:  []byte(`{"extraDataPackagePath":"test.zzz", "configuration":{"locale":"en-US"}}`),
			expectedARN: UploadARN,
		},
		{
			name:        "Both ARN and path missed, then no ARN should be presented",
			jsonString:  []byte(`{}`),
			expectedARN: "",
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			conf := createRun(v.jsonString, client)
			fmt.Println(conf.String())
			s := structs.New(conf)
			if v.expectedARN == "" {
				f, ok := s.FieldOk("Configuration")
				if ok && !f.IsZero() {
					f, ok = f.FieldOk("ExtraDataPackageArn")
					if ok && !f.IsZero() {
						t.Fatalf("Test:%s\n ExtraDataPackageArn shouldn't be presented if it's not expected", v.name)
					}
				}
			} else {
				f, ok := s.Field("Configuration").FieldOk("ExtraDataPackageArn")
				if !ok && f.IsZero() && v.expectedARN != "" {
					t.Fatalf("Test:%s\n ExtraDataPackageArn field isn't exist or has default value", v.name)
				}
				if v.expectedARN == "" && ok && !f.IsZero() {
					t.Fatalf("Test:%s\n ExtraDataPackageArn shouldn't be presented if it's not expected", v.name)
				}
				if v.expectedARN != "" && *conf.Configuration.ExtraDataPackageArn != v.expectedARN {
					t.Fatalf("Test:%s\n Expected: %s, actual: %s", v.name, v.expectedARN, *conf.Test.TestPackageArn)
				}
			}
		})
	}
}

func TestUploadAuxiliaryApps(t *testing.T) {
	server := CreateFakeServer()
	defer server.Close()

	client := &MockClient{
		UploadTest: true,
		FakeServer: server,
	}

	cases := []struct {
		name       string
		jsonString []byte
		qtyARN     int
	}{
		{
			name:       "If auxiliaryAppsPath is not specified, then nothing should be uploaded",
			jsonString: []byte(`{"name": "string"}`),
			qtyARN:     0,
		},
		{
			name:       "If auxiliaryAppsPath specified, then all items should be uploaded",
			jsonString: []byte(`{"auxiliaryAppsPath": [ "test.zzz", "test.zzz" ]}`),
			qtyARN:     2,
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			conf := createRun(v.jsonString, client)
			fmt.Println(conf.String())
			if getAuxiliaryAppSize(conf) != v.qtyARN {
				t.Fatalf("Should be %d auxiliary app ARN, actual number: %d", v.qtyARN, len(conf.Configuration.AuxiliaryApps))
			}
		})
	}
}

func createRun(bytes []byte, client devicefarmiface.DeviceFarmAPI) *devicefarm.ScheduleRunInput {
	p := &DeviceFarmRun{
		Client:  client,
		Config:  model.Transform(bytes),
		Project: "232323",
	}
	return createScheduleRunInput(p)
}

func getAuxiliaryAppSize(config *devicefarm.ScheduleRunInput) int {
	s := structs.New(config)
	fmt.Println("getAuxiliaryAppSize size:", len(s.Fields()))
	fmt.Println("getAuxiliaryAppSize fields:")
	for _, v := range s.Fields() {
		fmt.Println(v.Name())
	}
	f, ok := s.FieldOk("Configuration")
	if !ok || (ok && f.IsZero()) {
		return 0
	}
	found := false
	for _, v := range f.Fields() {
		if v.Name() == "AuxiliaryApps" {
			found = true
			break
		}
	}
	if !found {
		return 0
	}
	return len(config.Configuration.AuxiliaryApps)
}

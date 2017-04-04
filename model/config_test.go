package model

import "testing"

func TestTransformSuccess(t *testing.T) {
	config := createSmallTestData()
	if config.RunName != "name" {
		t.Error("Can't transform JSON to struct")
	}
	if config.Test.Filter != "" {
		t.Error("JSON transformation to struct is incorrect")
	}
}

func TestSizeOfArrays(t *testing.T) {
	config := createSmallTestData()
	if len(config.AdditionalData.AuxiliaryApps) != 0 {
		t.Error("Size of AdditionalData.AuxiliaryApps should be 0")
	}
	if len(config.Test.Parameters) != 0 {
		t.Error("Size of Test.Parameters should be 0")
	}
}

func TestContainMapData(t *testing.T) {
	config := createMapTestData()
	if len(config.Test.Parameters) != 1 {
		t.Error("Size of Test.Parameters should be 1")
	}
	if config.Test.Parameters["key"] != "value" {
		t.Error("Test.Parameters['key'] should return 'value'")
	}
}

func TestDontContainMapData(t *testing.T) {
	config := createSmallTestData()
	if len(config.Test.Parameters) != 0 {
		t.Error("Size of Test.Parameters should be 0")
	}
}

func TestGetUploadUnexisted(t *testing.T) {
	res := GetUploadTypeForTest("UIAUTOMATION")
	if res != "" {
		t.Error("For unexisted type of input result should be blank string")
	}
}

func TestGetUploadPositive(t *testing.T) {
	res := GetUploadTypeForTest("APPIUM_WEB_JAVA_JUNIT")
	if res != "APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE" {
		t.Error("For unexisted type of input result should be blank string")
	}
}

func TestExecutionConfiguration(t *testing.T) {
	config := createMapTestData()
	if config.ExecutionConfiguration.JobTimeoutMinutes != 0 {
		t.Error("By default timeout should be 0")
	}
}

func createSmallTestData() RunConfig {
	json := []byte(`{"runName":"name"}`)
	return Transform(json)
}

func createMapTestData() RunConfig {
	json := []byte(`{"runName":"name","test":{"parameters":{"key":"value"}}}`)
	return Transform(json)
}

package model

import (
	"testing"

	"github.com/fatih/structs"
)

func TestTransformDefault(t *testing.T) {
	config := createSmallTestData()
	if config.Name != "name" {
		t.Error("Can't transform JSON to struct")
	}
	s := structs.New(config)
	if !s.Field("Test").IsZero() {
		t.Error("JSON transformation to struct is incorrect")
	}
	if !s.Field("Configuration").IsZero() {
		t.Error("JSON transformation to struct is incorrect")
	}
	if !s.Field("ExecutionConfiguration").IsZero() {
		t.Error("JSON transformation to struct is incorrect")
	}
}

func TestTransformNonDefault(t *testing.T) {
	config := createMapTestData()
	if config.Name != "name" {
		t.Error("Can't transform JSON to struct")
	}
	if *config.ExecutionConfiguration.JobTimeoutMinutes != 11 {
		t.Error("Can't properly initialize config")
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

func createSmallTestData() *RunConfig {
	json := []byte(`{"name":"name"}`)
	return Transform(json)
}

func createMapTestData() *RunConfig {
	json := []byte(`{"name":"name","executionConfiguration":{"jobTimeoutMinutes":11,"accountsCleanup":true,"appPackagesCleanup":true}}`)
	return Transform(json)
}

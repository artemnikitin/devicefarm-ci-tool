package model

import (
	"encoding/json"
	"log"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

var m = map[string]string{
	"APPIUM_JAVA_JUNIT":      "APPIUM_JAVA_JUNIT_TEST_PACKAGE",
	"APPIUM_JAVA_TESTNG":     "APPIUM_JAVA_TESTNG_TEST_PACKAGE",
	"APPIUM_PYTHON":          "APPIUM_PYTHON_TEST_PACKAGE",
	"CALABASH":               "CALABASH_TEST_PACKAGE",
	"INSTRUMENTATION":        "INSTRUMENTATION_TEST_PACKAGE",
	"UIAUTOMATOR":            "UIAUTOMATOR_TEST_PACKAGE",
	"XCTEST":                 "XCTEST_TEST_PACKAGE",
	"XCTEST_UI":              "XCTEST_UI_TEST_PACKAGE",
	"APPIUM_WEB_JAVA_JUNIT":  "APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE",
	"APPIUM_WEB_JAVA_TESTNG": "APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE",
	"APPIUM_WEB_PYTHON":      "APPIUM_WEB_PYTHON_TEST_PACKAGE",
}

// RunConfig contains serialized representation of run model from JSON file
type RunConfig struct {
	Name                   string                               `json:"name,omitempty"`
	ProjectArn             string                               `json:"projectArn,omitempty"`
	ProjectName            string                               `json:"projectName,omitempty"`
	AppArn                 string                               `json:"appArn,omitempty"`
	AppPath                string                               `json:"appPath,omitempty"`
	DevicePoolArn          string                               `json:"devicePoolArn,omitempty"`
	DevicePoolName         string                               `json:"devicePoolName,omitempty"`
	TestPackagePath        string                               `json:"testPackagePath,omitempty"`
	ExtraDataPackagePath   string                               `json:"extraDataPackagePath,omitempty"`
	AuxiliaryAppsPath      []string                             `json:"auxiliaryAppsPath,omitempty"`
	Test                   *devicefarm.ScheduleRunTest          `json:"test,omitempty"`
	Configuration          *devicefarm.ScheduleRunConfiguration `json:"configuration,omitempty"`
	ExecutionConfiguration *devicefarm.ExecutionConfiguration   `json:"executionConfiguration,omitempty"`
}

// Transform unmarshall JSON model file to struct
func Transform(jsonBytes []byte) *RunConfig {
	result := &RunConfig{}
	err := json.Unmarshal(jsonBytes, result)
	errors.Validate(err, "Can't read model file")
	return result
}

// GetUploadTypeForTest return type of upload based on type of test
func GetUploadTypeForTest(testType string) string {
	v, exist := m[testType]
	if !exist {
		log.Println("Can't determine type of upload for", testType)
		return ""
	}
	return v
}

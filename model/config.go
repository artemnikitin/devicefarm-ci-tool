package model

import (
	"encoding/json"
	"log"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
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
	RunName string `json:"runName,omitempty"`
	Test    struct {
		Filter          string            `json:"filter,omitempty"`
		Parameters      map[string]string `json:"parameters,omitempty"`
		TestPackageArn  string            `json:"testPackageArn,omitempty"`
		TestPackagePath string            `json:"testPackagePath,omitempty"`
		Type            string            `json:"type,omitempty"`
	} `json:"test,omitempty"`
	AdditionalData struct {
		AuxiliaryApps        []string `json:"auxiliaryApps,omitempty"`
		BillingMethod        string   `json:"billingMethod,omitempty"`
		ExtraDataPackageArn  string   `json:"extraDataPackageArn,omitempty"`
		ExtraDataPackagePath string   `json:"extraDataPackagePath,omitempty"`
		Locale               string   `json:"locale,omitempty"`
		Location             struct {
			Latitude  float64 `json:"latitude,omitempty"`
			Longitude float64 `json:"longitude,omitempty"`
		} `json:"location,omitempty"`
		NetworkProfileArn string `json:"networkProfileArn,omitempty"`
		Radios            struct {
			Bluetooth string `json:"bluetooth,omitempty"`
			Gps       string `json:"gps,omitempty"`
			Nfc       string `json:"nfc,omitempty"`
			Wifi      string `json:"wifi,omitempty"`
		} `json:"radios,omitempty"`
	} `json:"additionalData,omitempty"`
}

// Transform unmarshall JSON model file to struct
func Transform(jsonBytes []byte) RunConfig {
	result := &RunConfig{}
	err := json.Unmarshal(jsonBytes, result)
	errors.Validate(err, "Can't read model file")
	return *result
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

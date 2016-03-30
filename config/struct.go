package config

import (
	"encoding/json"
	"log"

	"github.com/artemnikitin/devicefarm-ci-tool/errors"
)

const m = map[string]string{
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

// RunConfig contains serialized representation of run config from JSON file
type RunConfig struct {
	RunName string `json:"runName"`
	Test    struct {
		Filter          string            `json:"filter"`
		Parameters      map[string]string `json:"parameters"`
		TestPackageArn  string            `json:"testPackageArn"`
		TestPackagePath string            `json:"testPackagePath"`
		Type            string            `json:"type"`
	} `json:"test"`
	AdditionalData struct {
		AuxiliaryApps        []string `json:"auxiliaryApps"`
		BillingMethod        string   `json:"billingMethod"`
		ExtraDataPackageArn  string   `json:"extraDataPackageArn"`
		ExtraDataPackagePath string   `json:"extraDataPackagePath"`
		Locale               string   `json:"locale"`
		Location             struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		NetworkProfileArn string `json:"networkProfileArn"`
		Radios            struct {
			Bluetooth string `json:"bluetooth"`
			Gps       string `json:"gps"`
			Nfc       string `json:"nfc"`
			Wifi      string `json:"wifi"`
		} `json:"radios"`
	} `json:"additionalData"`
}

// Transform unmarshall JSON config file to struct
func Transform(jsonBytes []byte) RunConfig {
	result := &RunConfig{}
	err := json.Unmarshal(jsonBytes, result)
	errors.Validate(err, "Can't read config file")
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

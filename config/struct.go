package config

import (
	"encoding/json"
	"log"
)

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
	if err != nil {
		log.Fatal("Can't transform JSON to struct because of:", err)
	}
	return *result
}

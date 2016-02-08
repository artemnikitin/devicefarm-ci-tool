package config

// RunConfig contains serialized representation of run config from JSON file
type RunConfig struct {
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
			Bluetooth bool `json:"bluetooth"`
			Gps       bool `json:"gps"`
			Nfc       bool `json:"nfc"`
			Wifi      bool `json:"wifi"`
		} `json:"radios"`
	} `json:"additionalData"`
	RunName string `json:"runName"`
	Test    struct {
		Filter     string `json:"filter"`
		Parameters struct {
			String1 string `json:"string1"`
			String2 string `json:"string2"`
		} `json:"parameters"`
		TestPackageArn  string `json:"testPackageArn"`
		TestPackagePath string `json:"testPackagePath"`
		Type            string `json:"type"`
	} `json:"test"`
}

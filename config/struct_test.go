package config

import (
	"testing"
)

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
		t.Error("Size should be 0")
	}
	if len(config.Test.Parameters) != 0 {
		t.Error("Size should be 0")
	}
}

func TestContainMapData(t *testing.T) {
	config := createMapTestData()
	if len(config.Test.Parameters) != 1 {
		t.Error("Size of map should be 1")
	}
	if config.Test.Parameters["key"] != "value" {
		t.Error("Result should be equal to 'value'")
	}
}

func TestDontContainMapData(t *testing.T) {
	config := createSmallTestData()
	if len(config.Test.Parameters) != 0 {
		t.Error("Size of map should be 0")
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
package up

import (
	"encoding/json"
	"os"
	"testing"
)

const ConfigErrorMessage = "Got error from config read. Error: %v"
const skipIntTestMessage = "Skipping integration tests, set environment variable INTEGRATION."

type TestConfig struct {
	Token     string `json:"token"`
	AccountId string `json:"accountId"`
}

func GetTestConfig(fileName string) (*TestConfig, error) {
	var config TestConfig
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func SkipIfNotIntegrationTest(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip(skipIntTestMessage)
	}
}

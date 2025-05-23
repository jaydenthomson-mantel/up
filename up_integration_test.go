package up

import (
	"encoding/json"
	"os"
	"testing"
)

const configErrorMessage = "Got error from config read. Error: %v"
const skipIntTestMessage = "Skipping integration tests, set environment variable INTEGRATION."

type TestConfig struct {
	Token     string `json:"token"`
	AccountId string `json:"accountId"`
}

func TestGetAccounts(t *testing.T) {
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf(configErrorMessage, err)
		return
	}

	upClient := NewClient()

	firstAccount, err := upClient.GetAccounts(config.Token)
	if err != nil {
		t.Errorf("Got error from GetAccounts function. Error: %v", err)
		return
	}

	if firstAccount.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}
}

func TestGetTransactions(t *testing.T) {
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf(configErrorMessage, err)
		return
	}

	upClient := NewClient()

	transaction, err := upClient.GetTransactions(config.AccountId, config.Token)
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
	}

	if transaction.Data[0].ID == "" {
		t.Errorf("Id for first transaction fetched is empty.")
	}
}

func SkipIfNotIntegrationTest(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		t.Skip(skipIntTestMessage)
	}
}

func GetTestConfig() (*TestConfig, error) {
	var config TestConfig
	fileContent, err := os.ReadFile("./.config/testing.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

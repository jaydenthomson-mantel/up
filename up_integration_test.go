//go:build integration

package up

import (
	"encoding/json"
	"os"
	"testing"
)

type TestConfig struct {
	Token     string `json:"token"`
	AccountId string `json:"accountId"`
}

func TestGetAccounts(t *testing.T) {
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf("Got error from config read. Error: %v", err)
		return
	}

	upClient := NewClient()

	accounts, err := upClient.GetAccounts(config.Token, &PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
		return
	}

	if accounts.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}

	nextAccount, err := GetNextPage(upClient, (*PagedData[Account])(accounts), config.Token)
	if err != nil {
		t.Errorf("Got error from next page function. Error: %v", err)
	}

	if nextAccount.Data[0].ID == "" {
		t.Errorf("Id for second account fetched is empty.")
	}

	if accounts.Data[0].ID == nextAccount.Data[0].ID {
		t.Errorf("Id for first and second account fetched matched.")
	}
}

func TestGetTransactions(t *testing.T) {
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf("Got error from config read. Error: %v", err)
		return
	}

	upClient := NewClient()

	transaction, err := upClient.GetTransactions(config.AccountId, config.Token, &PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
	}

	if transaction.Data[0].ID == "" {
		t.Errorf("Id for first transaction fetched is empty.")
	}

	nextTransaction, err := GetNextPage(upClient, (*PagedData[Transaction])(transaction), config.Token)
	if err != nil {
		t.Errorf("Got error from next page function. Error: %v", err)
	}

	if nextTransaction.Data[0].ID == "" {
		t.Errorf("Id for second transaction fetched is empty.")
	}

	if transaction.Data[0].ID == nextTransaction.Data[0].ID {
		t.Errorf("Id for first and second transaction fetched matched.")
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

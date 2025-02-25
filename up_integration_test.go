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

	firstAccount, err := upClient.GetAccounts(config.Token, &PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
		return
	}

	if firstAccount.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}

	nextAccount, err := (*PagedData[Account])(firstAccount).GetNextPage(upClient, config.Token)
	if err != nil {
		t.Errorf("Got error from next page function. Error: %v", err)
	}

	if nextAccount.Data[0].ID == "" {
		t.Errorf("Id for second account fetched is empty.")
	}

	if firstAccount.Data[0].ID == nextAccount.Data[0].ID {
		t.Errorf("Id for first and second account fetched matched.")
	}

	accounts, err := (*PagedData[Account])(firstAccount).GetAllPages(upClient, config.Token)
	if err != nil {
		t.Errorf("Got error fetching all pages. Error: %v", err)
	}

	if accounts[0].Data[0].ID == "" || accounts[1].Data[0].ID == "" {
		t.Errorf("Some of the ID's fetched are empty when fetching all pages. Error: %v", err)
	}

	if accounts[0].Data[0].ID == accounts[1].Data[0].ID {
		t.Errorf("First ID and second ID are the same when fetching all pages. Error: %v", err)
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

	nextTransaction, err := (*PagedData[Transaction])(transaction).GetNextPage(upClient, config.Token)
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

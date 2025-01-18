package up

import (
	"encoding/json"
	"fmt"
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
	accounts, err := upClient.GetAccounts(config.Token, map[string]string{"page[size]": "1"})

	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
		return
	}

	jsonAccounts, _ := json.MarshalIndent(accounts, "", "  ")
	fmt.Println(string(jsonAccounts))
}

func TestGetTransactions(t *testing.T) {
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf("Got error from config read. Error: %v", err)
		return
	}

	upClient := NewClient()
	transactions, err := upClient.GetTransactions(config.AccountId, config.Token, map[string]string{"page[size]": "1"})

	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
	} else {
		jsonTransactions, _ := json.MarshalIndent(transactions, "", "  ")
		fmt.Println(string(jsonTransactions))
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

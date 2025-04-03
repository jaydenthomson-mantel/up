package up

import (
	"strconv"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig("./.config/testing.json")
	if err != nil {
		t.Errorf(ConfigErrorMessage, err)
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

	nextTransaction, err := upClient.GetNextTransactions(transaction, config.Token)
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

func TestGetTransactionMax(t *testing.T) {
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig("./.config/testing.json")
	if err != nil {
		t.Errorf(ConfigErrorMessage, err)
		return
	}

	upClient := NewClient()

	maxPageTransactions, err := upClient.GetTransactionMaxPage(config.AccountId, config.Token)
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
		return
	}

	if maxPageTransactions.Data[0].ID == "" {
		t.Errorf("Id for first transaction fetched is empty.")
	}

	maxPageSizeConversion, err := strconv.Atoi(MaxPageSize)
	if err != nil {
		t.Errorf("Got error from maxPageSize conversion. Error: %v", err)
	}

	pageLength := len(maxPageTransactions.Data)

	if pageLength != maxPageSizeConversion {
		t.Errorf(
			"Expected GetTransactionMaxPage to be equal to '%v' but was actually '%v'."+
				"Troubleshoot to ensure there are enough transactions to query.",
			maxPageSizeConversion,
			pageLength)
	}
}

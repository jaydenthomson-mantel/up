package transaction

import (
	"strconv"
	"testing"

	"github.com/jaydenthomson-mantel/up"
	"github.com/jaydenthomson-mantel/up/pagination"
)

func TestGetTransactions(t *testing.T) {
	up.SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := up.GetTestConfig("../.config/testing.json")
	if err != nil {
		t.Errorf(up.ConfigErrorMessage, err)
		return
	}

	upClient := up.NewClient()

	transaction, err := GetTransactions(upClient, config.AccountId, config.Token, &pagination.PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
	}

	if transaction.Data[0].ID == "" {
		t.Errorf("Id for first transaction fetched is empty.")
	}

	nextTransaction, err := GetNextTransactions(upClient, transaction, config.Token)
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
	up.SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := up.GetTestConfig("../.config/testing.json")
	if err != nil {
		t.Errorf(up.ConfigErrorMessage, err)
		return
	}

	upClient := up.NewClient()

	maxPageTransactions, err := GetTransactionMaxPage(upClient, config.AccountId, config.Token)
	if err != nil {
		t.Errorf("Got error from function. Error: %v", err)
		return
	}

	if maxPageTransactions.Data[0].ID == "" {
		t.Errorf("Id for first transaction fetched is empty.")
	}

	maxPageSizeConversion, err := strconv.Atoi(pagination.MaxPageSize)
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

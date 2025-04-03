package account

import (
	"testing"

	"github.com/jaydenthomson-mantel/up"
)

func TestGetAccounts(t *testing.T) {
	up.SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := up.GetTestConfig("../.config/testing.json")
	if err != nil {
		t.Errorf(up.ConfigErrorMessage, err)
		return
	}

	upClient := up.NewClient()

	firstAccount, err := GetAccounts(upClient, config.Token, &up.PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from GetAccounts function. Error: %v", err)
		return
	}

	if firstAccount.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}

	nextAccount, err := GetNextAccounts(upClient, firstAccount, config.Token)
	if err != nil {
		t.Errorf("Got error from next page function. Error: %v", err)
	}

	if nextAccount.Data[0].ID == "" {
		t.Errorf("Id for second account fetched is empty.")
	}

	if firstAccount.Data[0].ID == nextAccount.Data[0].ID {
		t.Errorf("Id for first and second account fetched matched.")
	}
}

func TestGetAllAccounts(t *testing.T) {
	up.SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := up.GetTestConfig("../.config/testing.json")
	if err != nil {
		t.Errorf(up.ConfigErrorMessage, err)
		return
	}

	upClient := up.NewClient()

	_, err = GetAllAccounts(upClient, config.Token)
	if err != nil {
		t.Errorf("Got error fetching all pages. Error: %v", err)
	}

	accountsMaxPage, err := GetAccountsMaxPage(upClient, config.Token)
	if err != nil {
		t.Errorf("Got error from GetAccounts function. Error: %v", err)
		return
	}

	if accountsMaxPage.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}
}

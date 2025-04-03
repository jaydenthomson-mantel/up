package up

import "testing"

func TestGetAccounts(t *testing.T) {
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf(configErrorMessage, err)
		return
	}

	upClient := NewClient()

	firstAccount, err := upClient.GetAccounts(config.Token, &PaginationParams{PageSize: "1"})
	if err != nil {
		t.Errorf("Got error from GetAccounts function. Error: %v", err)
		return
	}

	if firstAccount.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}

	nextAccount, err := upClient.GetNextAccounts(firstAccount, config.Token)
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
	SkipIfNotIntegrationTest(t)
	t.Parallel()
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf(configErrorMessage, err)
		return
	}

	upClient := NewClient()

	_, err = upClient.GetAllAccounts(config.Token)
	if err != nil {
		t.Errorf("Got error fetching all pages. Error: %v", err)
	}

	accountsMaxPage, err := upClient.GetAccountsMaxPage(config.Token)
	if err != nil {
		t.Errorf("Got error from GetAccounts function. Error: %v", err)
		return
	}

	if accountsMaxPage.Data[0].ID == "" {
		t.Errorf("Id for first account fetched is empty.")
	}
}

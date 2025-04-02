package up

import (
	"fmt"
	"time"
)

type Account struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		DisplayName   string `json:"displayName"`
		AccountType   string `json:"accountType"`
		OwnershipType string `json:"ownershipType"`
		Balance       struct {
			CurrencyCode     string `json:"currencyCode"`
			Value            string `json:"value"`
			ValueInBaseUnits int    `json:"valueInBaseUnits"`
		} `json:"balance"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"attributes"`
	Relationships struct {
		Transactions struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"transactions"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type AccountsResponse PagedData[Account]

func (up *UpClient) GetAccounts(token string, params *PaginationParams) (*AccountsResponse, error) {
	url := fmt.Sprintf("%v/accounts", up.baseUrl)
	return get[AccountsResponse](up, url, token, params)
}

func (up *UpClient) GetNextAccounts(accountResponse *AccountsResponse, token string) (*AccountsResponse, error) {
	nextAccounts, err := (*PagedData[Account])(accountResponse).GetNextPage(up, token)
	return (*AccountsResponse)(nextAccounts), err
}

func (up *UpClient) GetAllAccounts(token string) ([]*AccountsResponse, error) {
	accounts, err := up.GetAccountsMaxPage(token)
	if err != nil {
		return nil, err
	}

	allAccounts, err := (*PagedData[Account])(accounts).GetAllPages(up, token)
	if err != nil {
		return nil, err
	}

	result := make([]*AccountsResponse, len(allAccounts))
	result[0] = accounts
	for i, page := range allAccounts {
		result[i] = (*AccountsResponse)(page)
	}

	return result, nil
}

func (up *UpClient) GetAccountsMaxPage(token string) (*AccountsResponse, error) {
	params := &PaginationParams{PageSize: maxPageSize}
	return up.GetAccounts(token, params)
}

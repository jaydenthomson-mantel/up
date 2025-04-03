package account

import (
	"fmt"
	"time"

	"github.com/jaydenthomson-mantel/up"
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

type AccountsResponse up.PagedData[Account]

func GetAccounts(upClient *up.UpClient, token string, params *up.PaginationParams) (*AccountsResponse, error) {
	url := fmt.Sprintf("%v/accounts", upClient.BaseUrl)
	return up.Get[AccountsResponse](upClient, url, token, params)
}

func GetNextAccounts(upClient *up.UpClient, accountResponse *AccountsResponse, token string) (*AccountsResponse, error) {
	nextAccounts, err := (*up.PagedData[Account])(accountResponse).GetNextPage(upClient, token)
	return (*AccountsResponse)(nextAccounts), err
}

func GetAccountsMaxPage(upClient *up.UpClient, token string) (*AccountsResponse, error) {
	params := &up.PaginationParams{PageSize: up.MaxPageSize}
	return GetAccounts(upClient, token, params)
}

func GetAllAccounts(upClient *up.UpClient, token string) ([]*AccountsResponse, error) {
	accounts, err := GetAccountsMaxPage(upClient, token)
	if err != nil {
		return nil, err
	}

	allAccounts, err := (*up.PagedData[Account])(accounts).GetAllPages(upClient, token)
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

package account

import (
	"fmt"

	"github.com/jaydenthomson-mantel/up"
	"github.com/jaydenthomson-mantel/up/pagination"
)

func GetAccounts(upClient *up.UpClient, token string, params *pagination.PaginationParams) (*AccountsResponse, error) {
	url := fmt.Sprintf("%v/accounts", upClient.BaseUrl)
	return up.Get[AccountsResponse](upClient, url, token, params)
}

func GetNextAccounts(upClient *up.UpClient, accountResponse *AccountsResponse, token string) (*AccountsResponse, error) {
	nextAccounts, err := (*pagination.PagedData[Account])(accountResponse).GetNextPage(upClient, token)
	return (*AccountsResponse)(nextAccounts), err
}

func GetAccountsMaxPage(upClient *up.UpClient, token string) (*AccountsResponse, error) {
	params := &pagination.PaginationParams{PageSize: pagination.MaxPageSize}
	return GetAccounts(upClient, token, params)
}

func GetAllAccounts(upClient *up.UpClient, token string) ([]*AccountsResponse, error) {
	accounts, err := GetAccountsMaxPage(upClient, token)
	if err != nil {
		return nil, err
	}

	allAccounts, err := (*pagination.PagedData[Account])(accounts).GetAllPages(upClient, token)
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

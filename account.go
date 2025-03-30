package up

import (
	"fmt"

	"github.com/jaydenthomson-mantel/up/models"
)

func (up *UpClient) GetAccounts(token string, params *PaginationParams) (*models.AccountsResponse, error) {
	url := fmt.Sprintf("%v/accounts", up.baseUrl)
	return getAndUnmarshal[models.AccountsResponse](up, url, token, params)
}

func (up *UpClient) GetAccountsMaxPage(token string) (*models.AccountsResponse, error) {
	params := &PaginationParams{PageSize: maxPageSize}
	return up.GetAccounts(token, params)
}

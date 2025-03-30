package up

import (
	"fmt"

	"github.com/jaydenthomson-mantel/up/models"
)

func (up *UpClient) GetTransactions(accountId string, token string, params *PaginationParams) (*models.TransactionsResponse, error) {
	url := fmt.Sprintf("%v/accounts/%v/transactions", up.baseUrl, accountId)
	return getAndUnmarshal[models.TransactionsResponse](up, url, token, params)
}

func (up *UpClient) GetTransactionMaxPage(accountId string, token string) (*models.TransactionsResponse, error) {
	params := &PaginationParams{PageSize: maxPageSize}
	return up.GetTransactions(accountId, token, params)
}

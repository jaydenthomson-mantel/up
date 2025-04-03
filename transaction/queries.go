package transaction

import (
	"fmt"

	"github.com/jaydenthomson-mantel/up"
	"github.com/jaydenthomson-mantel/up/pagination"
)

func GetTransactions(upClient *up.UpClient, accountId string, token string, params *pagination.PaginationParams) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%v/accounts/%v/transactions", upClient.BaseUrl, accountId)
	return up.Get[TransactionsResponse](upClient, url, token, params)
}

func GetNextTransactions(upClient *up.UpClient, transactionResponse *TransactionsResponse, token string) (*TransactionsResponse, error) {
	nextTransactions, err := (*pagination.PagedData[Transaction])(transactionResponse).GetNextPage(upClient, token)
	return (*TransactionsResponse)(nextTransactions), err
}

func GetTransactionMaxPage(upClient *up.UpClient, accountId string, token string) (*TransactionsResponse, error) {
	params := &pagination.PaginationParams{PageSize: pagination.MaxPageSize}
	return GetTransactions(upClient, accountId, token, params)
}

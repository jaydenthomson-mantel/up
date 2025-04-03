package transaction

import (
	"fmt"
	"time"

	"github.com/jaydenthomson-mantel/up"
)

type Transaction struct {
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

type TransactionsResponse up.PagedData[Transaction]

func GetTransactions(upClient *up.UpClient, accountId string, token string, params *up.PaginationParams) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%v/accounts/%v/transactions", upClient.BaseUrl, accountId)
	return up.Get[TransactionsResponse](upClient, url, token, params)
}

func GetNextTransactions(upClient *up.UpClient, transactionResponse *TransactionsResponse, token string) (*TransactionsResponse, error) {
	nextTransactions, err := (*up.PagedData[Transaction])(transactionResponse).GetNextPage(upClient, token)
	return (*TransactionsResponse)(nextTransactions), err
}

func GetTransactionMaxPage(upClient *up.UpClient, accountId string, token string) (*TransactionsResponse, error) {
	params := &up.PaginationParams{PageSize: up.MaxPageSize}
	return GetTransactions(upClient, accountId, token, params)
}

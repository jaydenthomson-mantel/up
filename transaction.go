package up

import (
	"fmt"
	"time"
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

type TransactionsResponse PagedData[Transaction]

func (up *UpClient) GetTransactions(accountId string, token string, params *PaginationParams) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%v/accounts/%v/transactions", up.BaseUrl, accountId)
	return Get[TransactionsResponse](up, url, token, params)
}

func (up *UpClient) GetNextTransactions(transactionResponse *TransactionsResponse, token string) (*TransactionsResponse, error) {
	nextTransactions, err := (*PagedData[Transaction])(transactionResponse).GetNextPage(up, token)
	return (*TransactionsResponse)(nextTransactions), err
}

func (up *UpClient) GetTransactionMaxPage(accountId string, token string) (*TransactionsResponse, error) {
	params := &PaginationParams{PageSize: MaxPageSize}
	return up.GetTransactions(accountId, token, params)
}

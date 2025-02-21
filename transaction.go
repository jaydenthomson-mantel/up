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
	url := fmt.Sprintf("%v/accounts/%v/transactions", up.baseUrl, accountId)
	return get[TransactionsResponse](up, url, token, params)
}

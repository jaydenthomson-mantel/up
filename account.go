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
	var accountsResp AccountsResponse
	err := get(up, url, token, params, &accountsResp)
	if err != nil {
		return nil, err
	}

	return &accountsResp, nil
}

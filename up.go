package up

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type UpClient struct {
	httpClient http.Client
	baseUrl    string
}

// TODO make this a struct
type UpParams map[string]string

func NewClient() *UpClient {
	return &UpClient{
		httpClient: http.Client{
			Timeout: time.Second * 5,
		},
		baseUrl: "https://api.up.com.au/api/v1",
	}
}

func (up *UpClient) GetAccounts(token string, params UpParams) (*AccountsResponse, error) {
	url := fmt.Sprintf("%v/accounts", up.baseUrl)
	var accountsResp AccountsResponse
	err := get(up, url, token, params, &accountsResp)
	if err != nil {
		return nil, err
	}

	return &accountsResp, nil
}

func (up *UpClient) GetTransactions(accountId string, token string, params UpParams) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%v/accounts/%v/transactions", up.baseUrl, accountId)
	var transactionsResp TransactionsResponse
	err := get(up, url, token, params, &transactionsResp)
	if err != nil {
		return nil, err
	}

	return &transactionsResp, nil
}

func get[T any](up *UpClient, url string, token string, params UpParams, t *T) error {
	err := validateParams(token, params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	addParams(req, token, params)
	resp, err := up.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, t)
	if err != nil {
		return err
	}

	return nil
}

func validateParams(token string, params UpParams) error {
	pageSizeStr, hit := params["page[size]"]
	if hit {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > 100 {
			return PageSizeError{badPageSize: pageSizeStr}
		}
	}

	if false {
		// TODO add validation for token format
	}

	return nil
}

func addParams(req *http.Request, token string, params UpParams) {
	req.Header.Add("Authorization", "Bearer "+token)
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

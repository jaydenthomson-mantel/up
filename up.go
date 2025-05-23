package up

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jaydenthomson-mantel/up/models"
)

type UpClient struct {
	httpClient http.Client
	baseUrl    string
}

func NewClient() *UpClient {
	return &UpClient{
		httpClient: http.Client{
			Timeout: time.Second * 5,
		},
		baseUrl: "https://api.up.com.au/api/v1",
	}
}

func (up *UpClient) GetAccounts(token string) (*models.PageAccount, error) {
	url := fmt.Sprintf("%v/accounts", up.baseUrl)
	return get[models.PageAccount](up, url, token)
}

func (up *UpClient) GetAccount(accountId string, token string) (*models.AccountRecord, error) {
	url := fmt.Sprintf("%v/accounts/%v", up.baseUrl, accountId)
	return get[models.AccountRecord](up, url, token)
}

func (up *UpClient) GetTransactions(accountId *string, token string) (*models.TransactionPage, error) {
	var url string
	if accountId == nil {
		url = fmt.Sprintf("%v/transactions", up.baseUrl)
	} else {
		url = fmt.Sprintf("%v/accounts/%v/transactions", up.baseUrl, *accountId)
	}
	return get[models.TransactionPage](up, url, token)
}

func (up *UpClient) GetTransaction(transactionId string, token string) (*models.TransactionRecord, error) {
	url := fmt.Sprintf("%v/transactions/%v", up.baseUrl, transactionId)
	return get[models.TransactionRecord](up, url, token)
}

func get[T any](up *UpClient, url string, token string) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := up.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var t T
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

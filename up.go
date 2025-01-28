package up

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type UpClient struct {
	httpClient http.Client
	baseUrl    string
}

type QueryParams interface {
	Validate() error
	ToMap() map[string]string
}

func NewClient() *UpClient {
	return &UpClient{
		httpClient: http.Client{
			Timeout: time.Second * 5,
		},
		baseUrl: "https://api.up.com.au/api/v1",
	}
}

func get[T any](up *UpClient, url string, token string, params QueryParams, t *T) error {
	err := validate(token, params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	addToRequest(req, token, params)
	resp, err := up.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, t)
}

func validate(token string, params QueryParams) error {
	err := validateToken(token)
	if err != nil {
		return err
	}

	if params != nil {
		return params.Validate()
	}

	return nil
}

func addToRequest(req *http.Request, token string, params QueryParams) {
	req.Header.Add("Authorization", "Bearer "+token)
	if params != nil {
		q := req.URL.Query()
		m := params.ToMap()
		for key, value := range m {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}
}

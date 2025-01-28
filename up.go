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
	err = getResponse(up, req, t)
	if err != nil {
		return err
	}

	return nil
}

func getResponse[T any](up *UpClient, req *http.Request, t *T) error {
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

func validate(token string, params QueryParams) error {
	err := validateToken(token)
	if err != nil {
		return err
	}

	err = params.Validate()
	if err != nil {
		return err
	}

	return nil
}

func addToRequest(req *http.Request, token string, params QueryParams) {
	req.Header.Add("Authorization", "Bearer "+token)
	q := req.URL.Query()
	m := params.ToMap()
	for key, value := range m {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
}

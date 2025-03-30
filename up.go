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

// get makes HTTP GET requests to the Up API
func (up *UpClient) get(url string, token string, params QueryParams) ([]byte, error) {
	err := validate(token, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addToRequest(req, token, params)
	resp, err := up.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// GetPaged implements the common.Getter interface for pagination
func (up *UpClient) GetPaged(url string, token string) (interface{}, error) {
	result, err := getAndUnmarshal[map[string]interface{}](up, url, token, nil)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// getAndUnmarshal is a package-level generic function for making typed API requests
func getAndUnmarshal[T any](up *UpClient, url string, token string, params QueryParams) (*T, error) {
	body, err := up.get(url, token, params)
	if err != nil {
		return nil, err
	}

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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

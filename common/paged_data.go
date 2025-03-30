package common

import "encoding/json"

// PagedData represents a paginated response from the Up API
type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

// Getter defines the interface for making paginated requests
type Getter interface {
	GetPaged(url string, token string) (interface{}, error)
}

func (page *PagedData[T]) GetNextPage(up Getter, token string) (*PagedData[T], error) {
	url := page.Links.Next
	if url == "" {
		return nil, nil
	}

	resp, err := up.GetPaged(url, token)
	if err != nil {
		return nil, err
	}

	if rawData, ok := resp.(map[string]interface{}); ok {
		jsonData, err := json.Marshal(rawData)
		if err != nil {
			return nil, err
		}
		var pagedData PagedData[T]
		if err := json.Unmarshal(jsonData, &pagedData); err != nil {
			return nil, err
		}
		return &pagedData, nil
	}
	return nil, nil
}

func (page *PagedData[T]) GetAllPages(up Getter, token string) ([]*PagedData[T], error) {
	nextPageUrl := page.Links.Next
	pageList := []*PagedData[T]{page}

	for nextPageUrl != "" {
		resp, err := up.GetPaged(nextPageUrl, token)
		if err != nil {
			return nil, err
		}

		if rawData, ok := resp.(map[string]interface{}); ok {
			jsonData, err := json.Marshal(rawData)
			if err != nil {
				return nil, err
			}
			var nextPage PagedData[T]
			if err := json.Unmarshal(jsonData, &nextPage); err != nil {
				return nil, err
			}
			pageList = append(pageList, &nextPage)
			nextPageUrl = nextPage.Links.Next
		} else {
			break
		}
	}

	return pageList, nil
}

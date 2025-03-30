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
	return page.getPageFromURL(up, page.Links.Next, token)
}

func (page *PagedData[T]) GetAllPages(up Getter, token string) ([]*PagedData[T], error) {
	pageList := []*PagedData[T]{page}
	nextPage := page

	for nextPage.Links.Next != "" {
		var err error
		nextPage, err = nextPage.getPageFromURL(up, nextPage.Links.Next, token)
		if err != nil {
			return nil, err
		}
		if nextPage == nil {
			break
		}
		pageList = append(pageList, nextPage)
	}

	return pageList, nil
}

func (page *PagedData[T]) getPageFromURL(up Getter, url string, token string) (*PagedData[T], error) {
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

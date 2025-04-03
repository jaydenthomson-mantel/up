package pagination

import (
	"github.com/jaydenthomson-mantel/up"
)

const MaxPageSize = "100"

type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

func (page *PagedData[T]) GetNextPage(upClient *up.UpClient, token string) (*PagedData[T], error) {
	url := page.Links.Next
	if url == "" {
		return nil, nil
	}

	return up.Get[PagedData[T]](upClient, url, token, nil)
}

func (page *PagedData[T]) GetAllPages(upClient *up.UpClient, token string) ([]*PagedData[T], error) {
	nextPageUrl := page.Links.Next
	pageList := []*PagedData[T]{page}

	for nextPageUrl != "" {
		nextPage, err := up.Get[PagedData[T]](upClient, nextPageUrl, token, nil)
		if err != nil {
			return nil, err
		}
		pageList = append(pageList, nextPage)
		nextPageUrl = nextPage.Links.Next
	}

	return pageList, nil
}

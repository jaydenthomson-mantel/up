package up

import (
	"fmt"
	"strconv"
)

const maxPageSize = "100"

type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

type PaginationParams struct {
	PageSize string
}

type PageSizeError struct {
	BadPageSize string
}

func (page *PagedData[T]) GetNextPage(up *UpClient, token string) (*PagedData[T], error) {
	url := page.Links.Next
	if url == "" {
		return nil, nil
	}

	return get[PagedData[T]](up, url, token, nil)
}

func (page *PagedData[T]) GetAllPages(up *UpClient, token string) ([]*PagedData[T], error) {
	nextPageUrl := page.Links.Next
	pageList := []*PagedData[T]{page}

	for nextPageUrl != "" {
		nextPage, err := get[PagedData[T]](up, nextPageUrl, token, nil)
		if err != nil {
			return nil, err
		}
		pageList = append(pageList, nextPage)
		nextPageUrl = nextPage.Links.Next
	}

	return pageList, nil
}

func (params PaginationParams) Validate() error {
	pageSizeStr := params.PageSize
	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > 100 {
			return PageSizeError{BadPageSize: pageSizeStr}
		}
	}

	return nil
}

func (params PaginationParams) ToMap() map[string]string {
	m := make(map[string]string)
	if params.PageSize != "" {
		m["page[size]"] = params.PageSize
	}
	return m
}

func (p PageSizeError) Error() string {
	return fmt.Sprintf("page size %v not allowed", p.BadPageSize)
}

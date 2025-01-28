package up

import (
	"fmt"
	"strconv"
)

type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	}
}

type PaginationParams struct {
	PageSize string
}

type PageSizeError struct {
	BadPageSize string
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

func GetNextPage[T any](up *UpClient, page *PagedData[T], token string) (*PagedData[T], error) {
	url := page.Links.Next
	if url == "" {
		return nil, nil
	}

	var nextPage PagedData[T]
	err := get(up, url, token, nil, &nextPage)
	if err != nil {
		return nil, err
	}

	return &nextPage, err
}

func (p PageSizeError) Error() string {
	return fmt.Sprintf("page size %v not allowed", p.BadPageSize)
}

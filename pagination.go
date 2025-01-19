package up

import (
	"fmt"
	"strconv"
)

type PaginationParams struct {
	pageSize string
}

type PageSizeError struct {
	badPageSize string
}

func (params PaginationParams) Validate() error {
	pageSizeStr := params.pageSize
	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > 100 {
			return PageSizeError{badPageSize: pageSizeStr}
		}
	}

	return nil
}

func (params PaginationParams) ToMap() map[string]string {
	m := make(map[string]string)
	if params.pageSize != "" {
		m["page[size]"] = params.pageSize
	}
	return m
}

func (p PageSizeError) Error() string {
	return fmt.Sprintf("page size %v not allowed", p.badPageSize)
}

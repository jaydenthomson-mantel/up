package up

import (
	"strconv"

	"github.com/jaydenthomson-mantel/up/errors"
)

const maxPageSize = "100"

type PaginationParams struct {
	PageSize string
}

func (params PaginationParams) Validate() error {
	pageSizeStr := params.PageSize
	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 || pageSize > 100 {
			return errors.PageSizeError{BadPageSize: pageSizeStr}
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

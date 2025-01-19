package up

import "strconv"

type PaginationParams struct {
	pageSize string
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

package up

type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
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

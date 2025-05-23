package models

type RecordData[T any] struct {
	Data  T `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type PagedData[T any] struct {
	Data  []T `json:"data"`
	Links struct {
		Prev string `json:"prev"`
		Next string `json:"next"`
	} `json:"links"`
}

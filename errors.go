package up

import "fmt"

type PageSizeError struct {
	badPageSize string
}

func (p PageSizeError) Error() string {
	return fmt.Sprintf("page size %v not allowed", p.badPageSize)
}

type BearerFormatError struct{}

func (p BearerFormatError) Error() string {
	return "the bearer token was in an unknown format"
}

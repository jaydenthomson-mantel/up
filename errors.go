package up

import "fmt"

type PageSizeError struct {
	BadPageSize string
}

func (p PageSizeError) Error() string {
	return fmt.Sprintf("page size %v not allowed", p.BadPageSize)
}

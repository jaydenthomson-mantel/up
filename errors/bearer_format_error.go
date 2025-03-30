package errors

type BearerFormatError struct{}

func (p BearerFormatError) Error() string {
	return "the bearer token was in an unknown format"
}

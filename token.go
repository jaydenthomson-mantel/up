package up

import (
	"regexp"
)

type BearerFormatError struct{}

func (p BearerFormatError) Error() string {
	return "the bearer token was in an unknown format"
}

func validateToken(token string) error {
	match, err := regexp.MatchString("^up:yeah:[a-zA-Z0-9]+$", token)
	if err != nil {
		return err
	}
	if !match {
		return BearerFormatError{}
	}

	return nil
}

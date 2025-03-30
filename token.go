package up

import (
	"regexp"

	"github.com/jaydenthomson-mantel/up/errors"
)

func validateToken(token string) error {
	match, err := regexp.MatchString("^up:yeah:[a-zA-Z0-9]+$", token)
	if err != nil {
		return err
	}
	if !match {
		return errors.BearerFormatError{}
	}

	return nil
}

package errors

import (
	stderr "errors"
)

func Is(err, target error) bool {
	return stderr.Is(err, target)
}

func As(err error, target interface{}) bool {
	return stderr.As(err, target)
}

func Unwrap(err error) error {
	return stderr.Unwrap(err)
}

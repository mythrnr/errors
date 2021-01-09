package errors

import (
	stderr "errors"
)

type wrappedError struct {
	err   error
	cause error
}

var _ error = (*wrappedError)(nil)

func Wrap(err, cause error) error {
	if err == nil {
		return nil
	}

	return &wrappedError{
		err:   err,
		cause: cause,
	}
}

func (e *wrappedError) As(target interface{}) bool {
	return stderr.As(e.err, target) || stderr.As(e.cause, target)
}

func (e *wrappedError) Error() string {
	return e.err.Error()
}

func (e *wrappedError) Is(err error) bool {
	return stderr.Is(e.err, err) || stderr.Is(e.cause, err)
}

func (e *wrappedError) Unwrap() error {
	return e.cause
}

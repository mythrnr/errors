package errors

import "errors"

// As just calls standard `errors.As`.
//
// As は標準の `errors.As` を呼び出すだけ.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is just calls standard `errors.Is`.
//
// Is は標準の `errors.Is` を呼び出すだけ.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// New just calls standard `errors.New`.
//
// New は標準の `errors.New` を呼び出すだけ.
func New(text string) error {
	//nolint:err113
	return errors.New(text)
}

// Unwrap just calls standard `errors.Unwrap`.
//
// Unwrap は標準の `errors.Unwrap` を呼び出すだけ.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

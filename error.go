package errors

type stdErr interface {
	error
	As(interface{}) bool
	Is(error) bool
	Unwrap() error
}

type wrappingErr struct {
	main  error
	cause error
}

var _ stdErr = (*wrappingErr)(nil)

// Wrap returns the error object includes `err` and `cause` object.
// Returns `nil` if `err` is `nil`.
//
// Wrap は `err` と `cause` を含んだエラーオブジェクトを返す.
// `err` が `nil` の場合は `nil` が返される.
func Wrap(err, cause error) error {
	if err == nil {
		return nil
	}

	return &wrappingErr{
		main:  err,
		cause: cause,
	}
}

// As searches for value that can be assigned to `target`
// with the priority of `e.main`, `e.cause`,
// and returns `true` if the value can be assigned.
//
// As は `target` に代入可能な値を `e.main` , `e.cause` の優先順位で探し,
// 代入できた場合は `true` を返す.
func (e *wrappingErr) As(target interface{}) bool {
	return As(e.main, target) || As(e.cause, target)
}

// Error calls `Error` in `e.main` and returns message.
//
// Error は `e.main` の `Error` を呼び出してメッセージを返す.
func (e *wrappingErr) Error() string {
	return e.main.Error()
}

// Is searches for an error object matching `err`
// with the priority of `e.main`, `e.cause`,
// and returns `true` if it matches.
//
// Is は `err` に一致するエラーオブジェクトを
// `e.main` , `e.cause` の優先順位で探し, 一致した場合は `true` を返す.
func (e *wrappingErr) Is(err error) bool {
	return Is(e.main, err) || Is(e.cause, err)
}

// Unwrap returns the inner `e.cause`.
//
// Unwrap は内側の `e.cause` を返す.
func (e *wrappingErr) Unwrap() error {
	return e.cause
}

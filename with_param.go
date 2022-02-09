package errors

// WithParamsError is used when you want to add a parameter
// to the error to handle it.
// The type is exposed so that it can be determined using `errors.As`.
//
// WithParamsError はエラーにパラメータを追加してハンドリングしたい場合に使う.
// `errors.As` で判定できるように型は公開しておく.
type WithParamsError struct {
	err error
	ps  []interface{}
}

var _ stdError = (*WithParamsError)(nil)

// NewWithParamsError creates a new `WithParamsError`.
// If the `err` in the argument is `nil`, `nil` is returned.
//
// NewWithParamsError は新規の `WithParamsError` を生成する.
// 引数の `err` が `nil` の場合は `nil` が返される.
func NewWithParamsError(err error, ps ...interface{}) *WithParamsError {
	if err == nil {
		return nil
	}

	return &WithParamsError{err: err, ps: ps}
}

// As returns `true` and assign `e.err` if the `target` can be assigned.
//
// As は `e.err` が `target` に代入可能であれば代入して `true` を返す.
func (e *WithParamsError) As(target interface{}) bool {
	return As(e.err, target)
}

// Error returns the result of calling `Error` in `e.err`.
//
// Error は `e.err` の `Error` を呼び出して返す.
func (e *WithParamsError) Error() string {
	return e.err.Error()
}

// Is returns whether `err` matches `e.err` or not.
//
// Is は `err` が `e.err` に一致するかどうかを返す.
func (e *WithParamsError) Is(err error) bool {
	return Is(e.err, err)
}

// Unwrap returns the inner `e.err`.
//
// Unwrap は内側の `e.err` を返す.
func (e *WithParamsError) Unwrap() error {
	return e.err
}

// Params returns the parameters given at creation time.
//
// Params は生成時に与えられたパラメータを返す.
func (e *WithParamsError) Params() []interface{} {
	return e.ps
}

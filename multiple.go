package errors

import "strings"

// MultipleError is used when you want to handle multiple errors.
// The type is exposed so that it can be determined using `errors.As`.
//
// MultipleError は複数のエラーをハンドリングしたい場合に使う.
// `errors.As` で判定できるように型は公開しておく.
type MultipleError struct{ errs []error }

var _ stdJoinError = (*MultipleError)(nil)

// NewMultipleError creates a new `MultipleError`.
// The error in the argument that is `nil` is excluded.
// If non-nil error is `0`, `nil` is returned.
//
// NewMultipleError は新規の `MultipleError` を生成する.
// 引数の error の内, `nil` のものは除外される.
// `nil` でない error が `0` の場合, `nil` が返される.
func NewMultipleError(errs ...error) *MultipleError {
	m := &MultipleError{errs: make([]error, 0, len(errs))}

	for _, err := range errs {
		if err != nil {
			m.errs = append(m.errs, err)
		}
	}

	if len(m.errs) == 0 {
		return nil
	}

	return m
}

// As returns `true` after assigning to `target` the first element in
// `m.errs` that is assignable to it.
//
// As は `m.errs` の中に `target` に代入可能な最初の要素があれば代入して `true` を返す.
func (m *MultipleError) As(target interface{}) bool {
	for _, err := range m.errs {
		if As(err, target) {
			return true
		}
	}

	return false
}

// Error calls `Error` of the enclosing error
// and returns a comma-separated concatenated string.
//
// Error は内包する error の `Error` を呼び出し,
// カンマ区切りで連結した文字列を返す.
func (m *MultipleError) Error() string {
	var buf strings.Builder

	for _, err := range m.errs {
		buf.WriteString(err.Error())
		buf.WriteByte(',')
	}

	return buf.String()[:buf.Len()-1]
}

// Is returns whether there is an element in `m.errs` that matches `err`.
//
// Is は `m.errs` の中に `err` に一致する要素があるかどうかを返す.
func (m *MultipleError) Is(err error) bool {
	for _, e := range m.errs {
		if Is(e, err) {
			return true
		}
	}

	return false
}

// Unwrap returns `m.errs`.
//
// Unwrap は `m.errs` を返す.
func (m *MultipleError) Unwrap() []error { return m.errs }

// Errs returns a slice of the enclosing error.
//
// Errs は内包する error のスライスを返す.
func (m *MultipleError) Errs() []error { return m.Unwrap() }

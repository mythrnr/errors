package errors

import "strings"

// Multiple is used when you want to handle multiple errors.
// The type is exposed so that it can be determined using `errors.As`.
//
// Multiple は複数のエラーをハンドリングしたい場合に使う.
// `errors.As` で判定できるように型は公開しておく.
type Multiple struct{ errs []error }

var _ error = (*Multiple)(nil)

// NewMultiple creates a new `Multiple`.
// The error in the argument that is `nil` is excluded.
// If non-nil error is `0`, `nil` is returned.
//
// NewMultiple は新規の `Multiple` を生成する.
// 引数の error の内, `nil` のものは除外される.
// `nil` でない error が `0` の場合, `nil` が返される.
func NewMultiple(errs ...error) *Multiple {
	m := &Multiple{errs: make([]error, 0, len(errs))}

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

// Error calls `Error` of the enclosing error
// and returns a comma-separated concatenated string.
//
// Error は内包する error の `Error` を呼び出し,
// カンマ区切りで連結した文字列を返す.
func (m *Multiple) Error() string {
	var buf strings.Builder

	for _, err := range m.errs {
		buf.WriteString(err.Error())
		buf.WriteByte(',')
	}

	return buf.String()[:buf.Len()-1]
}

// Errs returns a slice of the enclosing error.
//
// Errs は内包する error のスライスを返す.
func (m *Multiple) Errs() []error { return m.errs }

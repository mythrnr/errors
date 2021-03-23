// Package errors makes enable to wrap error object by error object.
//
// By using `Is`, `As`, and `Unwrap` of the standard package `errors`,
// you can extract and use errors contained in them.
//
// - When wrapping, pass the error object as is instead of a string message
// like `fmt.Errorf("%w", err)`.
//
// - Since the error object is contained, predefined errors can be nested
// and judged by `errors.Is`.
//
// ---- In Japanese ----
//
// Package errors はエラーオブジェクトをエラーオブジェクトで包むことを可能にします.
//
// 標準パッケージである `errors` の `Is` , `As` , `Unwrap` を利用することで,
// 内包するエラーを取り出して利用することができる.
//
// - ラップするときに `fmt.Errorf("%w", err)` のように文字列のメッセージではなく
// エラーオブジェクトをそのまま渡す.
//
// - エラーオブジェクトが内包されているので, 定義済みエラーをネストさせて
// `errors.Is` で判定ができる.
package errors

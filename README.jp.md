# mythrnr/errors

[English](./README.md)

## Status

[![Check codes](https://github.com/mythrnr/errors/actions/workflows/check_code.yml/badge.svg)](https://github.com/mythrnr/errors/actions/workflows/check_code.yml)

[![Create Release](https://github.com/mythrnr/errors/actions/workflows/release.yml/badge.svg)](https://github.com/mythrnr/errors/actions/workflows/release.yml)

## Description

- `mythrnr/errors` は複数のエラーをより便利に取り扱う機能を提供する.  
- `Wrap` , `MultipleError` が追加されている以外は標準の `errors` パッケージと同じ.
- `errors.New` , `errors.Is` , `errors.As` , `errors.Unwrap` は標準パッケージを呼び出している.

## Features

### Wrapping error by error

- ラップするときに `fmt.Errorf("%w", err)` のように文字列のメッセージではなくエラーオブジェクトをそのまま渡す.
- エラーオブジェクトを内包させるので, 定義済みエラーをネストさせて `errors.Is` で判定ができる.
- `errors.Is` , `errors.As` , `errors.Unwrap` を利用することで,
内包するエラーを取り出して利用することができる.  

#### Problem

`fmt.Errorf` を使ったラップ処理では, 定義済みのエラー同士を階層化できない.

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

var (
    ErrCauseA = errors.New("error caused by A")
    ErrCauseB = errors.New("error caused by B")
)

func main() {
    err := fmt.Errorf("error caused by B: %w", ErrCauseA)

    // Of course true.
    fmt.Println(errors.Is(err, ErrCauseA))

    // Oh, how we check the error is same as ErrCauseB ?
    fmt.Println(strings.Contains(err.Error(), ErrCauseB.Error()))
}
```

#### Solves

`mythrnr/errors` を使うと下記の通り.

```go
package main

import "github.com/mythrnr/errors"

var (
    ErrCauseA = errors.New("error caused by A")
    ErrCauseB = errors.New("error caused by B")
)

func main() {
    err := errors.Wrap(ErrCauseB, ErrCauseA)

    // true
    fmt.Println(errors.Is(err, ErrCauseA))

    // true!
    fmt.Println(errors.Is(err, ErrCauseB))
}
```

### Mutiple errors

- 複数のエラーをまとめて返却したい場合に使用する.

```go
package main

import (
    "fmt"
    "strings"

    "github.com/mythrnr/errors"
)

func returnErrors() error {
    return errors.NewMultipleError(
        errors.New("error 1"),
        errors.New("error 2"),
        errors.New("error 3"),
    )
}

func main() {
    err := returnErrors()

    // output: error 1,error 2,error 3
    fmt.Println(err.Error())

    errs := &errors.MultipleError{}
    if errors.As(err, &errs) {
        // output: 3
        fmt.Println(len(errs.Errs()))

        // output: error 1
        fmt.Println(errs.Errs()[0])
    }
}
```

## Requirements

Go 1.16 以上で確認をしている.

## Install

`go get` で取得する.

```bash
go get github.com/mythrnr/errors
```

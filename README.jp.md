# mythrnr/errors

[English](./README.md)

## Status

[![Check codes](https://github.com/mythrnr/errors/workflows/Check%20codes/badge.svg)](https://github.com/mythrnr/errors/actions?query=workflow%3A%22Check+codes%22)

## Description

`mythrnr/errors` はエラーをラップする機能を提供する.  
Go言語の標準パッケージである `errors` の `Is` , `As` , `Unwrap` を利用することで,
内包するエラーを取り出して利用することができる.  
`mythrnr/errors` はシンプルな実装なので, もう誰かが実現しているかもしれない...

### Feature

- ラップするときに `fmt.Errorf("%w", err)` のように文字列のメッセージではなく `error` オブジェクトをそのまま渡す.
- `error` オブジェクトを内包させるので, 定義済みエラーをネストさせて `errors.Is` で判定ができる.

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

## Requirements

Go 1.13 以上で確認をしている.

## Install

`go get` で取得する.

```bash
go get github.com/mythrnr/errors
```

## Usage

`Wrap` が追加されている以外は標準の `errors` パッケージとほぼ同じ.

### `errors.New` , `errors.Is` , `errors.As` , `errors.Unwrap` について

[std_errors.go](https://github.com/mythrnr/errors/blob/master/std_errors.go) にある通り,
標準の `errors` パッケージの同名関数を呼び出している.

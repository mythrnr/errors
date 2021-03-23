# mythrnr/errors

[日本語](./README.jp.md)

## Status

[![Check codes](https://github.com/mythrnr/errors/workflows/Check%20codes/badge.svg)](https://github.com/mythrnr/errors/actions?query=workflow%3A%22Check+codes%22)

[![Create Release](https://github.com/mythrnr/errors/workflows/Create%20Release/badge.svg)](https://github.com/mythrnr/errors/actions?query=workflow%3A%22Create+Release%22)

## Description

Package `mythrnr/errors` makes enable to wrap error object by error object.  
By using `Is`, `As`, and `Unwrap` of the standard package `errors`,
you can extract and use errors contained in them.

### Feature

- When wrapping, pass the error object as is instead of a string message
  like `fmt.Errorf("%w", err)`.
- Since the error object is contained, predefined errors can be nested
  and judged by `errors.Is`.

#### Problem

The wrapping process using `fmt.Errorf` cannot hierarchize predefined errors.

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

If you use `mythrnr/errors`, you get the following.

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

Go 1.13 or above.

## Install

Get it with `go get`.

```bash
go get github.com/mythrnr/errors
```

## Usage

Almost the same as the standard `errors` package,
except for the addition of `Wrap`.

### About `errors.New` , `errors.Is` , `errors.As` , `errors.Unwrap`

As shown in [std_errors.go](https://github.com/mythrnr/errors/blob/master/std_errors.go),
it calls the homonymous function of the standard `errors` package.

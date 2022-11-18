# mythrnr/errors

[日本語](./README.jp.md)

## Status

[![Check codes](https://github.com/mythrnr/errors/actions/workflows/check-code.yaml/badge.svg)](https://github.com/mythrnr/errors/actions/workflows/check-code.yaml)

[![Create Release](https://github.com/mythrnr/errors/actions/workflows/release.yaml/badge.svg)](https://github.com/mythrnr/errors/actions/workflows/release.yaml)

[![Scan Vulnerabilities](https://github.com/mythrnr/errors/actions/workflows/scan-vulnerabilities.yaml/badge.svg)](https://github.com/mythrnr/errors/actions/workflows/scan-vulnerabilities.yaml)

## Description

- Package `mythrnr/errors` provides functions to treat multiple errors more useful.
- This package is the same as the standard `errors` package,
except that `Wrap` and `MultipleError` have been added.
- `errors.New` , `errors.Is` , `errors.As` , `errors.Unwrap` calls just standard `errors` .

## Features

### Wrapping error by error

- When wrapping, pass the error object as is instead of a string message
  like `fmt.Errorf("%w", err)`.
- Since the error object is contained, predefined errors can be nested
  and judged by `errors.Is`.
- You can use included errors by using `errors.Is` , `errors.As` , `errors.Unwrap` .

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

### Mutiple errors

- Using this if you want to return multiple errors at once.

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

Go 1.16 or above.

## Install

Get it with `go get`.

```bash
go get github.com/mythrnr/errors
```

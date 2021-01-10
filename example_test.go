package errors_test

import (
	"fmt"

	"github.com/mythrnr/errors"
)

var (
	ErrCauseA = errors.New("error caused by A")
	ErrCauseB = errors.New("error caused by B")
)

func Example() {
	err := errors.Wrap(ErrCauseB, ErrCauseA)

	// true
	fmt.Println(errors.Is(err, ErrCauseA))

	// true
	fmt.Println(errors.Is(err, ErrCauseB))
}

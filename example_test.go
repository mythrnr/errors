package errors_test

import (
	"fmt"
	"strings"

	"github.com/mythrnr/errors"
)

var (
	ErrCauseA = errors.New("error caused by A")
	ErrCauseB = errors.New("error caused by B")
)

func Example() {
	// Standard wrapping
	{
		err := fmt.Errorf("error caused by B: %w", ErrCauseA)

		// Of course true.
		fmt.Println(errors.Is(err, ErrCauseA))

		// Oh, how we check the error is same as ErrCauseB ?
		fmt.Println(strings.Contains(err.Error(), ErrCauseB.Error()))
	}

	// Solves
	{
		err := errors.Wrap(ErrCauseB, ErrCauseA)

		// true
		fmt.Println(errors.Is(err, ErrCauseA))

		// true!
		fmt.Println(errors.Is(err, ErrCauseB))
	}
}

func ExampleNewMultipleError() {
	errs := errors.NewMultipleError(
		errors.New("error A"),
		errors.New("error B"),
		errors.New("error C"),
	)

	// error A,error B,error C
	fmt.Println(errs.Error())
}

func ExampleNewMultipleError_nil() {
	// returns nil
	errs := errors.NewMultipleError(nil, nil)

	// <nil>
	fmt.Println(errs)

	// It panics.
	// fmt.Println(errs.Error())
}

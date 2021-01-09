package errors_test

import (
	stderr "errors"
	"fmt"
	"testing"

	"github.com/mythrnr/errors"
	"github.com/stretchr/testify/assert"
)

type myError struct{ msg string }

func (e *myError) Error() string { return e.msg }

func Test_Wrap(t *testing.T) {
	assert.Nil(t, errors.Wrap(nil, nil))
	assert.NotNil(t, errors.Wrap(fmt.Errorf("error"), nil))
	assert.NotNil(t, errors.Wrap(fmt.Errorf("error"), fmt.Errorf("cause")))
}

func Test_wrappedError_As(t *testing.T) {
	{
		err := errors.Wrap(fmt.Errorf("error"), nil)
		me := &myError{}

		assert.False(t, stderr.As(err, &me))
		assert.Empty(t, me.Error())
	}

	{
		err := errors.Wrap(fmt.Errorf("error"), &myError{msg: "cause"})
		me := &myError{}

		assert.True(t, stderr.As(err, &me))
		assert.Equal(t, "cause", me.Error())
	}
}

func Test_wrappedError_Error(t *testing.T) {
	assert.Equal(t,
		"error",
		errors.Wrap(fmt.Errorf("error"), nil).Error(),
	)

	assert.Equal(t,
		"error",
		errors.Wrap(fmt.Errorf("error"), fmt.Errorf("cause")).Error(),
	)
}

func Test_wrappedError_Is(t *testing.T) {
	ErrTest := fmt.Errorf("defined error")

	{
		err := errors.Wrap(fmt.Errorf("error"), nil)
		assert.False(t, stderr.Is(err, ErrTest))
	}

	{
		err := errors.Wrap(fmt.Errorf("error"), fmt.Errorf("cause"))
		assert.False(t, stderr.Is(err, ErrTest))
	}

	{
		err := errors.Wrap(fmt.Errorf("error"), fmt.Errorf("defined error"))
		assert.False(t, stderr.Is(err, ErrTest))
	}

	{
		err := errors.Wrap(ErrTest, fmt.Errorf("error"))
		assert.True(t, stderr.Is(err, ErrTest))
	}

	{
		err := errors.Wrap(fmt.Errorf("error"), ErrTest)
		assert.True(t, stderr.Is(err, ErrTest))
	}
}

func Test_wrappedError_Unwrap(t *testing.T) {
	{
		err := errors.Wrap(fmt.Errorf("error"), nil)
		assert.Nil(t, stderr.Unwrap(err))
	}

	{
		err := errors.Wrap(fmt.Errorf("error"), fmt.Errorf("cause"))
		ue := stderr.Unwrap(err)
		assert.NotNil(t, ue)
		assert.Equal(t, "cause", ue.Error())
	}
}

package errors_test

import (
	"testing"

	"github.com/mythrnr/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type myError struct{ msg string }

func (e *myError) Error() string { return e.msg }

func Test_Wrap(t *testing.T) {
	t.Parallel()

	require.NoError(t, errors.Wrap(nil, nil))
	require.Error(t, errors.Wrap(errors.New("main"), nil))
	assert.Error(t, errors.Wrap(errors.New("main"), errors.New("cause")))
}

func Test_wrappingError_As(t *testing.T) {
	t.Parallel()

	t.Run("Not includes myError", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), nil)
		me := &myError{}

		assert.False(t, errors.As(err, &me))
		assert.Empty(t, me.Error())
	})

	t.Run("Including myError", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), &myError{msg: "cause"})
		me := &myError{}

		assert.True(t, errors.As(err, &me))
		assert.Equal(t, "cause", me.Error())
	})

	t.Run("Top level matching", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(&myError{msg: "main"}, &myError{msg: "cause"})
		me := &myError{}

		assert.True(t, errors.As(err, &me))
		assert.Equal(t, "main", me.Error())
	})
}

func Test_wrappingError_Error(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		"main",
		errors.Wrap(errors.New("main"), nil).Error(),
	)

	assert.Equal(t,
		"main: cause",
		errors.Wrap(errors.New("main"), errors.New("cause")).Error(),
	)
}

func Test_wrappingError_Is(t *testing.T) {
	t.Parallel()

	ErrTest := errors.New("defined error")

	t.Run("Not match, includes nil", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), nil)
		assert.False(t, errors.Is(err, ErrTest))
	})

	t.Run("Not match", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), errors.New("cause"))
		assert.False(t, errors.Is(err, ErrTest))
	})

	t.Run("Not match, error has same message includes", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), errors.New("defined error"))
		assert.False(t, errors.Is(err, ErrTest))
	})

	t.Run("Match, top level", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(ErrTest, errors.New("cause"))
		assert.True(t, errors.Is(err, ErrTest))
	})

	t.Run("Match, wrapped error", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), ErrTest)
		assert.True(t, errors.Is(err, ErrTest))
	})
}

func Test_wrappingError_Unwrap(t *testing.T) {
	t.Parallel()

	t.Run("Not wrapped", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), nil)
		assert.NoError(t, errors.Unwrap(err))
	})

	t.Run("Wrapped", func(t *testing.T) {
		t.Parallel()

		err := errors.Wrap(errors.New("main"), errors.New("cause"))
		ue := errors.Unwrap(err)

		require.Error(t, ue)
		assert.Equal(t, "cause", ue.Error())
	})
}

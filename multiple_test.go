package errors_test

import (
	"testing"

	"github.com/mythrnr/errors"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func Test_MultipleError(t *testing.T) {
	t.Parallel()

	t.Run("No Error", func(t *testing.T) {
		t.Parallel()

		// NOTE: assert.NoError is not accept typed-nil.
		{
			assert.Nil(t, errors.NewMultipleError())
			assert.Nil(t, errors.NewMultipleError(nil))
			assert.Nil(t, errors.NewMultipleError(nil, nil))
		}
	})

	t.Run("Single Error", func(t *testing.T) {
		t.Parallel()

		another := &Err3{}
		e := &Err1{}
		m := errors.NewMultipleError(e)

		assert.Equal(t, "error 1", m.Error())
		assert.True(t, errors.Is(m, e))
		assert.False(t, errors.Is(m, another))

		var (
			_e       *Err1
			_another *Err3
		)

		assert.True(t, errors.As(m, &_e))
		assert.False(t, errors.As(m, &_another))

		assert.Equal(t, []error{e}, m.Unwrap())
		assert.Len(t, m.Errs(), 1)
	})

	t.Run("Multiple Error", func(t *testing.T) {
		t.Parallel()

		another := &Err3{}
		e1 := &Err1{}
		e2 := &Err2{}
		m := errors.NewMultipleError(e1, nil, e2)

		assert.Equal(t, "error 1,error 2", m.Error())
		assert.True(t, errors.Is(m, e1))
		assert.True(t, errors.Is(m, e2))
		assert.False(t, errors.Is(m, another))

		var (
			_e1      *Err1
			_e2      *Err2
			_another *Err3
		)

		assert.True(t, errors.As(m, &_e1))
		assert.True(t, errors.As(m, &_e2))
		assert.False(t, errors.As(m, &_another))

		assert.Equal(t, []error{e1, e2}, m.Unwrap())
		assert.Len(t, m.Errs(), 2)
	})
}

//nolint:errname
type Err1 struct{}

func (e *Err1) Error() string { return "error 1" }

//nolint:errname
type Err2 struct{}

func (e *Err2) Error() string { return "error 2" }

//nolint:errname
type Err3 struct{}

func (e *Err3) Error() string { return "error 3" }

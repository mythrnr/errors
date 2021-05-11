package errors_test

import (
	"testing"

	"github.com/mythrnr/errors"
	"github.com/stretchr/testify/assert"
)

func Test_Multiple(t *testing.T) {
	t.Parallel()

	t.Run("No Error", func(t *testing.T) {
		t.Parallel()

		assert.Nil(t, errors.NewMultiple())
		assert.Nil(t, errors.NewMultiple(nil))
		assert.Nil(t, errors.NewMultiple(nil, nil))
	})

	t.Run("Single Error", func(t *testing.T) {
		t.Parallel()

		m := errors.NewMultiple(errors.New("first error"))

		assert.Equal(t, "first error", m.Error())
	})

	t.Run("Multiple Error", func(t *testing.T) {
		t.Parallel()

		m := errors.NewMultiple(
			errors.New("first error"),
			nil,
			errors.New("second error"),
		)

		assert.Equal(t, "first error,second error", m.Error())
		assert.Equal(t, "first error", m.Errs()[0].Error())
		assert.Equal(t, "second error", m.Errs()[1].Error())
	})
}
package errors_test

import (
	"testing"

	"github.com/mythrnr/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WithParamsError(t *testing.T) {
	t.Parallel()

	assert.Nil(t, errors.NewWithParamsError(nil, nil))
	assert.Nil(t, errors.NewWithParamsError(nil, 1, false))

	expected := errors.New("err")
	err := errors.NewWithParamsError(expected, 1, false, "string")
	require.NotNil(t, err)

	err1 := &errors.WithParamsError{}
	assert.ErrorAs(t, err, &err1)
	assert.Equal(t, "err", err1.Error())

	err2 := errors.New("other")
	assert.True(t, err.As(&err2))
	assert.Equal(t, "err", err2.Error())

	assert.ErrorIs(t, err, expected)
	assert.ErrorIs(t, err.Unwrap(), expected)

	assert.Len(t, err.Params(), 3)
	assert.Equal(t, err.Params()[0], 1)
	assert.Equal(t, err.Params()[1], false)
	assert.Equal(t, err.Params()[2], "string")
}

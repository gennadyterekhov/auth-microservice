package input

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanInput(t *testing.T) {
	inputMock := NewMock()

	var i int64
	inputMock.AddInts(5)
	err := inputMock.ScanInts(&i)
	require.NoError(t, err)
	require.Equal(t, int64(5), i)

	var str string
	inputMock.AddStrings("value")

	err = inputMock.ScanStrings(&str)
	require.NoError(t, err)
	require.Equal(t, "value", str)

	var val1, val2 string
	inputMock.AddStrings("val1", "val2")

	err = inputMock.ScanStrings(&val1, &val2)
	require.NoError(t, err)
	require.Equal(t, "val1", val1)
	require.Equal(t, "val2", val2)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFor_SortCsvColumns_OK(t *testing.T) {
	actual := Decode("")
	expected := 0
	require.Equal(t, expected, actual)

	actual = Decode("I")
	expected = 1
	require.Equal(t, expected, actual)

	actual = Decode("XXI")
	expected = 21
	require.Equal(t, expected, actual)

	actual = Decode("IV")
	expected = 4
	require.Equal(t, expected, actual)
}

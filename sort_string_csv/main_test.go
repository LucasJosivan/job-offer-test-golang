package sort_string_csv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFor_SortCsvColumns_OK(t *testing.T) {
	actual := SortCsvColumns("Beth,Charles,Danielle,Adam,Eric\n17945,10091,10088,3907,10132\n2,12,13,48,11")
	expected := "Adam,Beth,Charles,Danielle,Eric\n3907,17945,10091,10088,10132\n48,2,12,13,11"
	require.Equal(t, expected, actual)

	actual = SortCsvColumns("Charles,Beth,Danielle,Adam,Eric\n10091,17945,10088,3907,10132\n12,2,13,48,11")
	expected = "Adam,Beth,Charles,Danielle,Eric\n3907,17945,10091,10088,10132\n48,2,12,13,11"
	require.Equal(t, expected, actual)
}

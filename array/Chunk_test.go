package array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunkIntArray(t *testing.T) {
	res, err := ChunkIntArray([]int{1, 2, 3, 4, 5}, 3)
	require.NoError(t, err)
	require.EqualValues(t, res, [][]int{{1, 2, 3}, {4, 5}})

	res, err = ChunkIntArray([]int{1, 2, 3, 4, 5}, 0)
	require.NoError(t, err)
	require.EqualValues(t, res, [][]int{{1, 2, 3, 4, 5}})

	_, err = ChunkIntArray([]int{}, 3)
	require.EqualError(t, err, "empty array has been provided")

	_, err = ChunkIntArray([]int{1, 2, 3, 4, 5}, 10)
	require.EqualError(t, err, "array size is bigger than the array length")

	_, err = ChunkIntArray([]int{1, 2, 3, 4, 5}, -1)
	require.EqualError(t, err, "chunk size is invalid")

}

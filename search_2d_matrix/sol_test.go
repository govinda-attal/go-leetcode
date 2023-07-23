package search2dmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5}}

	assert.False(t, SearchMatrix(matrix, 4))

	matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	assert.True(t, SearchMatrix(matrix, 3))

}

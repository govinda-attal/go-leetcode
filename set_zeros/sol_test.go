package set_zeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeros(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}

	SetZeroes(matrix)

	assert.Equal(t, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, matrix)

	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 0, 7, 8},
		{0, 10, 11, 12},
		{13, 14, 15, 0},
	}

	SetZeroes(matrix)

	assert.Equal(t, [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, matrix)

	matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}}

	SetZeroes(matrix)

	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 4, 5, 0},
		{0, 3, 1, 0}}, matrix)

}

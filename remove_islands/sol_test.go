package remove_islands_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/remove_islands"
	"github.com/stretchr/testify/assert"
)

func TestRemoveIslands(t *testing.T) {
	type testCase struct {
		grid     [][]int
		expected [][]int
	}

	cases := []testCase{
		{
			grid: [][]int{
				{1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 1},
				{0, 0, 1, 0, 1, 0},
				{1, 1, 0, 0, 1, 0},
				{1, 0, 1, 1, 0, 0},
				{1, 0, 0, 0, 0, 1},
			},
			expected: [][]int{
				{1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1, 0},
				{1, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 1},
			},
		},
		{
			grid: [][]int{
				{1, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 1},
			},
			expected: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 1},
			},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			out := RemoveIslands(tc.grid)
			assert.Equal(t, tc.expected, out)
		})
	}
}

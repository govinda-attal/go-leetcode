package shortest_path_binary_matrix_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/shortest_path_binary_matrix"
	"github.com/stretchr/testify/assert"
)

func TestShortedPathBinaryMatrix(t *testing.T) {
	type testCase struct {
		grid   [][]int
		result int
	}

	cases := []testCase{
		{
			grid:   [][]int{{0, 1}, {1, 0}},
			result: 2,
		},
		{
			grid:   [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			result: 4,
		},
		{
			grid:   [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			result: -1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, tc.result, ShortestPathBinaryMatrix(tc.grid))
		})
	}

}

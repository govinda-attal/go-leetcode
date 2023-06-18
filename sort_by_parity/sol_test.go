package sort_by_parity_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/sort_by_parity"

	"gotest.tools/assert"
)

func TestSortByParity(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
	}

	cases := []testCase{
		{
			input:    []int{3, 1, 2, 4},
			expected: []int{2, 4, 3, 1},
		},
		{
			input:    []int{0, 2, 1},
			expected: []int{0, 2, 1},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.DeepEqual(t, tc.expected, SortArrayByParity(tc.input))
		})
	}
}

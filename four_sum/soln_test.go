package four_sum_test

import (
	"fmt"
	"testing"

	foursum "github.com/govinda-attal/go-leetcode/four_sum"
	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	type testCase struct {
		nums     []int
		target   int
		expected [][]int
	}

	cases := []testCase{
		{
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:     []int{-3, -2, -1, 0, 0, 1, 2, 3},
			target:   0,
			expected: [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			out := foursum.FourSum(tc.nums, tc.target)
			assert.ElementsMatch(t, tc.expected, out)
		})
	}
}

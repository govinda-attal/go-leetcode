package max_subarray_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/max_sub_array"
	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	type testCase struct {
		input []int
		sum   int
	}

	cases := []testCase{
		{
			input: []int{-1, 3, 4, -5, 9, -2},
			sum:   11,
		},
		{
			input: []int{-2, -1},
			sum:   -1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, tc.sum, MaxSubArrayKA(tc.input))
		})
	}
}

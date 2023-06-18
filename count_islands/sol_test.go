package count_islands_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/count_islands"
	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	type testCase struct {
		grid  [][]byte
		count int
	}

	cases := []testCase{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			count: 1,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, tc.count, NumIslands(tc.grid))
		})
	}
}

package min_dist_farthest_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/min_dist_farthest"
	"github.com/stretchr/testify/assert"
)

func TestThis(t *testing.T) {
	type testCase struct {
		blocks    [][]string
		amenities []string
		outBlkNum int
	}
	tcs := []testCase{
		{
			blocks: [][]string{
				{"restaurant", "grocery"},
				{"cinema"},
				{"school"},
				{},
				{"school"},
			},
			amenities: []string{"school", "grocery"},
			outBlkNum: 1,
		},
		{
			blocks: [][]string{
				{"school"},
				{"gym"},
				{"gym", "school"},
				{"school"},
				{"school", "store"},
			},
			amenities: []string{"school", "gym", "store"},
			outBlkNum: 3,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, MinDistBlkToFarthestPoint(tc.blocks, tc.amenities), tc.outBlkNum)
		})
	}
}

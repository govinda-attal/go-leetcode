package three_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	exp := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	assert.ElementsMatch(t, exp, ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}

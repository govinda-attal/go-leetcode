package longest_consecutive_len

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, LongestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)

	assert.Equal(t, LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
}

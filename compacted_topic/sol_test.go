package compacted_topic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompact(t *testing.T) {
	assert.ElementsMatch(t, []int{5, 3, 2, 4, 1}, Compact([]int{1, 2, 3, 4, 5, 2, 4, 1}))
}

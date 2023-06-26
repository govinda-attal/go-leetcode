package grid_traveller_test

import (
	"testing"

	. "github.com/govinda-attal/go-leetcode/learnings/grid_traveller"
	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	assert.Equal(t, 1, NumberOfWays(1, 1))
	assert.Equal(t, 3, NumberOfWays(2, 3))
	assert.Equal(t, 6, NumberOfWays(3, 3))
}

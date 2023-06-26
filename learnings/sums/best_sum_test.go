package sums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestSum(t *testing.T) {
	assert.ElementsMatch(t, BestSum(7, []int{5, 3, 4}), []int{4, 3})
	assert.ElementsMatch(t, BestSum(7, []int{7, 3, 4}), []int{7})
	assert.ElementsMatch(t, BestSum(7, []int{3, 4, 7}), []int{7})
	assert.ElementsMatch(t, BestSum(7, []int{2, 4}), nil)
}

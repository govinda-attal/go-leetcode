package sums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHowSum(t *testing.T) {
	assert.ElementsMatch(t, HowSum(7, []int{5, 3, 4}), []int{4, 3})
	assert.ElementsMatch(t, HowSum(7, []int{2, 3}), []int{3, 2, 2})
	assert.ElementsMatch(t, HowSum(7, []int{2, 4}), nil)
}

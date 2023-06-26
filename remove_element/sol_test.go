package remove_element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	assert.Equal(t, 5, RemoveElement(arr, 2))
	// assert.ElementsMatch(t, []int{0, 1, 4, 0, 3}, arr)
}

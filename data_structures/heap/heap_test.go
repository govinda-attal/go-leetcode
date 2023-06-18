package heap_test

import (
	"testing"

	"github.com/govinda-attal/go-leetcode/data_structures/heap"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := heap.New[int](10)
	h.Add(10)
	h.Add(5)
	h.Add(2)
	h.Add(1)
	h.Add(7)
	h.Add(8)

	assert.Equal(t, []int{1, 2, 5, 10, 7, 8}, h.Items())

	val, _ := h.Poll()

	assert.Equal(t, 1, val)

	assert.Equal(t, []int{2, 7, 5, 10, 8}, h.Items())
}

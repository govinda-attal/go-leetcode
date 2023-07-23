package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8}, Fib(7))

	assert.Equal(t, 8, FibMemoized(6))
}

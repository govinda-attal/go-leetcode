package redundant_connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThis(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}

	assert.Equal(t, []int{2, 3}, FindRedundantConnection(edges))
}

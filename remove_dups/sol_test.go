package remove_dups

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDups(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	out := RemoveDuplicates(arr)
	assert.Equal(t, 5, out)
	expArray := []int{0, 1, 2, 3, 4}
	for i := 0; i < len(expArray); i++ {
		assert.Equal(t, expArray[i], arr[i])
	}

	fmt.Println(arr)
}

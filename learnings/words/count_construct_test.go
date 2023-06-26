package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountConstruct(t *testing.T) {
	assert.Equal(t, CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}), 1)
	assert.Equal(t, CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}), 2)
}

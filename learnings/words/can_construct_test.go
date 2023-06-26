package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}), true)
	assert.Equal(t, CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}), false)
}

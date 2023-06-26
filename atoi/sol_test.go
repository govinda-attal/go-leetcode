package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, -42, myAtoi("     -42"))
}

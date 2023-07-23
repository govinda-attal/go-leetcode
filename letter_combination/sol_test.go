package letter_combination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombination(t *testing.T) {
	// digits := "2"
	// exp := []string{"a", "b", "c"}
	// assert.ElementsMatch(t, exp, LetterCombinations(digits))

	digits := "23"
	exp := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.ElementsMatch(t, exp, LetterCombinations(digits))
}

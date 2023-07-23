package letter_combination

import (
	"fmt"
)

var numPad = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'0': " ",
}

func LetterCombinations(digits string) (rs []string) {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		val := numPad[digits[0]]
		rs = make([]string, len(val))
		for i, r := range val {
			rs[i] = string(r)
		}
		return rs
	}
	first := digits[0]

	restResult := LetterCombinations(digits[1:])
	for _, r := range numPad[first] {
		for _, s := range restResult {
			rs = append(rs, fmt.Sprintf("%c%s", r, s))
		}
	}
	return rs
}

package top_frequent_words_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/top_frequent_words"

	"github.com/stretchr/testify/assert"
)

func TestSortedListToBST(t *testing.T) {
	type testCase struct {
		input    []string
		count    int
		expected []string
	}

	cases := []testCase{
		{
			input:    []string{"i", "love", "leetcode", "i", "love", "coding"},
			count:    2,
			expected: []string{"i", "love"},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, TopKFrequent(tc.input, tc.count))
		})
	}
}

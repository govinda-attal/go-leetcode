package revlink_list_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/rev_linklist"
	"github.com/stretchr/testify/assert"

	ll "github.com/govinda-attal/go-leetcode/data_structures/link_list"
)

func TestSortedListToBST(t *testing.T) {
	type testCase struct {
		input    *ListNode
		expected *ListNode
	}
	tcs := []testCase{
		{input: ll.NewFromSlice[int](1, 2, 3), expected: ll.NewFromSlice[int](3, 2, 1)},
		{input: ll.NewFromSlice[int](1, 2, 3, 4), expected: ll.NewFromSlice[int](4, 3, 2, 1)},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected.ToSlice(), Reverse(tc.input).ToSlice())
		})
	}
}

package sorted_list_bst_test

import (
	"fmt"
	"testing"

	. "github.com/govinda-attal/go-leetcode/sorted_list_bst"

	ll "github.com/govinda-attal/go-leetcode/data_structures/link_list"
	"github.com/govinda-attal/go-leetcode/data_structures/tree"

	"gotest.tools/assert"
)

func TestSortedListToBST(t *testing.T) {
	type testCase struct {
		input    *ListNode
		expected *TreeNode
	}

	zero, minusThree, nine, minusTen, five := 0, -3, 9, -10, 5
	items := []*int{&zero, &minusThree, &nine, &minusTen, nil, &five, nil}

	cases := []testCase{
		{
			input:    ll.NewFromSlice[int](-10, -3, 0, 5, 9),
			expected: tree.NewFromBalancedSearchSlice[int](items),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			assert.DeepEqual(t, tc.expected.String(), SortedListToBST(tc.input).String())
		})
	}
}

package sorted_list_bst

import (
	ll "github.com/govinda-attal/go-leetcode/data_structures/link_list"
	"github.com/govinda-attal/go-leetcode/data_structures/tree"
)

// Problem Statement: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/

type TreeNode = tree.TreeNode[int]
type ListNode = ll.ListNode[int]

func SortedListToBST(head *ListNode) *TreeNode {
	items := head.ToSlice()
	return sortedArrayToBST(items)
}

// 1,2,3
func sortedArrayToBST(items []int) *TreeNode {
	if len(items) == 0 {
		return nil
	}
	if len(items) == 1 {
		return &TreeNode{
			Val: items[0],
		}
	}
	mid := len(items) / 2
	node := &TreeNode{
		Val:   items[mid],
		Left:  sortedArrayToBST(items[:mid]),
		Right: sortedArrayToBST(items[mid+1:]),
	}

	return node
}

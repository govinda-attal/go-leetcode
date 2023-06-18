package revlink_list

import "github.com/govinda-attal/go-leetcode/data_structures/link_list"

//https://leetcode.com/problems/reverse-linked-list/

type ListNode = link_list.ListNode[int]

// 1->2-3->nil
// 3->2->1->nil
func Reverse(head *ListNode) *ListNode {
	var prev, cur, next *ListNode = nil, head, head
	for {
		if cur == nil {
			return next
		}
		// nil | 1 | 2
		// tmp_prev := prev
		// 1 | 2 | 3
		next = cur
		// 2 | 3 | nil
		tmp_next := cur.Next

		// 1->nil | 2-> 1 | 3->2
		cur.Next = prev
		// 1 | 2 | 3
		prev = cur
		// 2 | 3 | nil
		cur = tmp_next
	}
}

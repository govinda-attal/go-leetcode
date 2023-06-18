package tree

import (
	"fmt"
	"strings"
)

type comparable interface {
	~int | ~float32 | ~int64 | ~float64 | ~int16 | ~int8
}

type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewFromBalancedSearchSlice[T comparable](items []*T) *TreeNode[T] {
	return newFromBinarySearchSlice[T](items, 0)
}

func newFromBinarySearchSlice[T comparable](items []*T, i int) *TreeNode[T] {
	if i >= len(items) {
		return nil
	}
	if items[i] == nil {
		return nil
	}
	l, r := 2*i+1, 2*i+2
	return &TreeNode[T]{
		Val:   *items[i],
		Left:  newFromBinarySearchSlice[T](items, l),
		Right: newFromBinarySearchSlice[T](items, r),
	}
}

func (tn *TreeNode[T]) ToSlice() []*T {
	rs := []*T{}
	queue := []*TreeNode[T]{tn}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {
			rs = append(rs, nil)
			continue
		}
		val := cur.Val
		rs = append(rs, &val)
		if cur.Left != nil {
			queue = append(queue, cur.Left, cur.Right)
		}
	}
	return rs
}

func (tn *TreeNode[T]) String() string {
	vals := []string{}
	items := tn.ToSlice()
	for _, item := range items {
		if item == nil {
			vals = append(vals, "nil")
		} else {
			vals = append(vals, fmt.Sprintf("%v", *item))
		}
	}
	return strings.Join(vals, ", ")
}

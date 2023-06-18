package tree_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/govinda-attal/go-leetcode/data_structures/tree"
	"github.com/stretchr/testify/assert"
)

type TreeNode = tree.TreeNode[int]

func TestBinarySearchTree(t *testing.T) {
	zero, minusThree, nine, minusTen, five := 0, -3, 9, -10, 5
	items := []*int{&zero, &minusThree, &nine, &minusTen, nil, &five, nil}
	toString := func(items []*int) string {
		vals := []string{}
		for _, item := range items {
			if item == nil {
				vals = append(vals, "nil")
			} else {
				vals = append(vals, fmt.Sprintf("%d", *item))
			}
		}
		return strings.Join(vals, ", ")
	}
	assert.EqualValues(t, toString(items), toString(tree.NewFromBalancedSearchSlice[int](items).ToSlice()))
}

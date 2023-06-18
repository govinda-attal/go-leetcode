package trie_test

import (
	"testing"

	"github.com/govinda-attal/go-leetcode/data_structures/trie"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {

	tree := trie.NewTrie("hello")
	tree.Add("hellopizza")

	count, ok := tree.FindWord("hello")
	assert.True(t, ok)
	assert.Equal(t, 5, count)
	count, ok = tree.FindWord("hellopizza")
	assert.True(t, ok)
	assert.Equal(t, 10, count)

	count, ok = tree.FindShortestWord("hellopizza")
	assert.True(t, ok)
	assert.Equal(t, 5, count)

	assert.ElementsMatch(t, []string{"hello", "hellopizza"}, tree.Words())
}

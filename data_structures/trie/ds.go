package trie

import "fmt"

type Node struct {
	val      rune
	ends     bool
	children map[rune]*Node
}

func NewTrie(word string) *Node {
	if len(word) == 0 {
		return nil
	}
	n := &Node{
		val: rune(word[0]),
	}
	if n.Add(word) {
		return n
	}
	return nil
}

func (n *Node) Add(word string) (ok bool) {
	size := len(word)
	if size == 0 {
		return false
	}
	first := rune(word[0])

	if n.val != first {
		return false
	}

	if size == 1 {
		n.ends = true
		return true
	}

	nextRune := rune(word[1])

	if n.children == nil {
		n.children = make(map[rune]*Node)
	}

	child, ok := n.children[nextRune]
	if !ok {
		child = &Node{
			val: nextRune,
		}
		n.children[nextRune] = child
	}

	return child.Add(word[1:])
}

func (n *Node) FindWord(word string) (size int, ok bool) {
	if len(word) == 0 {
		return 0, n == nil
	}
	if n.val != rune(word[0]) {
		return 0, false
	}
	if len(word) > 1 {
		child, ok := n.children[rune(word[1])]
		if !ok {
			return 0, false
		}
		size, ok = child.FindWord(word[1:])
		if !ok {
			return 0, false
		}
		return size + 1, true
	}
	return 1, true
}

func (n *Node) FindShortestWord(word string) (size int, ok bool) {
	if len(word) == 0 {
		return 0, n == nil
	}
	if n.val != rune(word[0]) {
		return 0, false
	}
	if n.ends {
		return 1, true
	}
	if len(word) > 1 {
		child, ok := n.children[rune(word[1])]
		if !ok {
			return 0, false
		}
		size, ok = child.FindShortestWord(word[1:])
		if !ok {
			return 0, false
		}
		return size + 1, true
	}
	return 1, true
}

func (n *Node) Words() []string {
	if n.children == nil {
		return []string{string(n.val)}
	}
	var rs []string
	if n.ends {
		rs = append(rs, string(n.val))
	}
	for _, c := range n.children {
		css := c.Words()
		for _, ss := range css {

			rs = append(rs, fmt.Sprintf("%c%s", n.val, ss))
		}
	}
	return rs
}

package adjacency_list_test

import (
	"strings"
	"testing"

	. "github.com/govinda-attal/go-leetcode/data_structures/adjacency_list"
	"github.com/stretchr/testify/assert"
)

func TestAdjacencyList(t *testing.T) {

	airports := strings.Split("BKK OKC JFK LAX MEX EZE HEL LOS LAP LIM", " ")

	routes := []Edge{
		{"PHX", "LAX"},
		{"PHX", "JFK"},
		{"JFK", "OKC"},
		{"JFK", "HEL"},
		{"JFK", "LOS"},
		{"MEX", "LAX"},
		{"MEX", "BKK"},
		{"MEX", "LIM"},
		{"MEX", "EZE"},
		{"LIM", "BKK"},
	}

	al := New()

	al.AddNodes(airports...)
	al.AddEdges(routes...)

	// t.Log(al)

	assert.True(t, al.BfsCanReach("PHX", "BKK"))
	assert.True(t, al.DfsCanReach("PHX", "BKK"))
}

package adjacency_list

import (
	"fmt"
	"strings"
)

type Node = string
type AdjacencyList map[string][]string
type Edge [2]Node

func New() AdjacencyList {
	return make(AdjacencyList)
}

func (al AdjacencyList) AddNodes(nn ...Node) {
	for _, n := range nn {
		al[n] = []string{}
	}
}

func (al AdjacencyList) AddEdges(ee ...Edge) {
	for _, e := range ee {
		al[e[0]] = append(al[e[0]], e[1])
		al[e[1]] = append(al[e[1]], e[0])
	}
}

func (al AdjacencyList) String() string {
	sb := strings.Builder{}
	for n, ee := range al {
		sb.WriteString(fmt.Sprintf("From [%s]: ", n) + fmt.Sprintf("%v\n", ee))
	}
	return sb.String()
}

func (al AdjacencyList) BfsCanReach(from, to string) bool {
	visited := map[Node]bool{}

	queue := []string{from}
	count := 0

	for len(queue) > 0 {
		count++
		next := queue[0]
		queue = queue[1:]
		if visited[next] {
			continue
		}
		if next == to {
			return true
		}
		for _, e := range al[next] {
			if visited[e] {
				continue
			}
			queue = append(queue, e)
		}
		visited[next] = true
	}
	return false
}

func (al AdjacencyList) DfsCanReach(from, to string) bool {
	if from == to {
		return true
	}

	visited := map[Node]bool{}
	// defer func() {
	// 	fmt.Println(visited)
	// }()
	return al.dfs(from, to, visited)
}

func (al AdjacencyList) dfs(from, to string, visited map[Node]bool) bool {
	if from == to {
		visited[to] = true
		return true
	}
	visited[from] = true
	for _, next := range al[from] {
		if next == to {
			visited[to] = true
			return true
		}
		if visited[next] {
			continue
		}
		if al.dfs(next, to, visited) {
			return true
		}
	}
	return false
}

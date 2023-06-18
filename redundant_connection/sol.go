package redundant_connection

func FindRedundantConnection(edges [][]int) []int {

	nodeNeigbours := make(map[int][]int)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _, ok := nodeNeigbours[a]; !ok {
			nodeNeigbours[a] = []int{}
		}
		if _, ok := nodeNeigbours[b]; !ok {
			nodeNeigbours[b] = []int{}
		}

		if dfsRedundantConnection(a, b, nodeNeigbours, make(map[int]bool)) {
			return edge
		}
		nodeNeigbours[a] = append(nodeNeigbours[a], b)
		nodeNeigbours[b] = append(nodeNeigbours[b], a)
	}
	return nil
}

func dfsRedundantConnection(src, target int, nodeNeigbours map[int][]int, visited map[int]bool) (rs bool) {
	if src == target {
		return true
	}

	visited[src] = true

	for _, neighbour := range nodeNeigbours[src] {
		if _, ok := visited[neighbour]; ok {
			continue
		}
		rs = rs || dfsRedundantConnection(neighbour, target, nodeNeigbours, visited)
	}
	return rs
}

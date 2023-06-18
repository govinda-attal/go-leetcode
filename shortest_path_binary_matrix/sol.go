package shortest_path_binary_matrix

// Problem Statement: https://leetcode.com/problems/shortest-path-in-binary-matrix/

type Point struct{ X, Y int }

var dirs = []Point{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func ShortestPathBinaryMatrix(grid [][]int) int {

	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	queue := []Point{{0, 0}}
	layer := 0
	grid[0][0] = 1
	for len(queue) != 0 {
		layer++
		num := len(queue)

		for i := 0; i < num; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.X == n-1 && node.Y == n-1 {
				return layer
			}

			for _, dir := range dirs {
				next := Point{node.X + dir.X, node.Y + dir.Y}
				if next.X < 0 || next.X >= n || next.Y < 0 || next.Y >= n || grid[next.X][next.Y] == 1 {
					continue
				}
				queue = append(queue, next)
				grid[next.X][next.Y] = 1
			}
		}
	}

	return -1
}

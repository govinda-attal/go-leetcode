package count_islands

// Problem Statement: https://leetcode.com/problems/number-of-islands/

func NumIslands(grid [][]byte) int {
	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				// dfs(grid, i, j)
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c int) {

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2'

	dfs(grid, r-1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
	dfs(grid, r+1, c)
}

type Cord = [2]int
type Queue = []Cord

func bfs(grid [][]byte, r, c int) {
	queue := Queue{Cord{r, c}}
	grid[r][c] = '2'

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		dirs := [4]Cord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, d := range dirs {
			x, y := cur[0]+d[0], cur[1]+d[1]

			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] != '1' {
				continue
			}
			queue = append(queue, Cord{x, y})
		}
	}
}

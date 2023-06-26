package grid_traveller

func NumberOfWays(m, n int) int {
	grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}
	grid[1][1] = 1

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			val := grid[i][j]
			if i+1 <= m {
				grid[i+1][j] += val
			}
			if j+1 <= n {
				grid[i][j+1] += val
			}
		}
	}
	return grid[m][n]
}

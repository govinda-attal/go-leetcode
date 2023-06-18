package remove_islands

// Problem Statement: https://www.youtube.com/watch?v=4tYoVx0QoN0
// https://www.tutorialspoint.com/program-to-remove-all-islands-which-are-completely-surrounded-by-water-in-python

func RemoveIslands(grid [][]int) [][]int {
	if len(grid) <= 2 || len(grid[0]) <= 2 {
		return grid
	}

	visited := visitedMatrix(len(grid), len(grid[0]))

	upLand := func(x, y int) bool {
		ux, uy := x, y-1
		return (uy >= 0) && rune(grid[ux][uy]) == 1
	}

	leftLand := func(x, y int) bool {
		lx, ly := x-1, y
		return (lx >= 0) && rune(grid[lx][ly]) == 1
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if rune(grid[i][j]) == 1 && !(upLand(i, j) || leftLand(i, j)) {
				if isIsland(grid, visited, i, j) {
					removeIsland(grid, i, j)
				}
			}
		}
	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	return grid
}

var dirs = map[rune][2]int{
	'l': {0, -1},
	'r': {0, 1},
	'u': {-1, 0},
	'd': {1, 0},
}

func isIsland(grid [][]int, visited [][]int, r, c int) bool {
	mr, mc := len(grid)-1, len(grid[0])-1

	visited[r][c] = 1

	beached := map[rune]bool{
		'l': false,
		'r': false,
		'u': false,
		'd': false,
	}
	for a, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if visited[nr][nc] != 0 {
			delete(beached, a)
			continue
		}

		if isOnEndge(nr, nc, mr, mc) {
			if grid[nr][nc] == 1 {
				return false
			}
			beached[a] = true
			continue
		}
		if grid[nr][nc] == 0 {
			visited[nr][nc] = 1
			beached[a] = true
			continue
		}
		beached[a] = isIsland(grid, visited, nr, nc)

	}

	if len(beached) == 0 {
		return false
	}

	rs := true
	for _, ok := range beached {
		rs = rs && ok
	}

	return rs
}

func removeIsland(grid [][]int, r, c int) {
	mr, mc := len(grid)-1, len(grid[0])-1
	grid[r][c] = 0
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if isOnEndge(nr, nc, mr, mc) {
			continue
		}
		if grid[nr][nc] == 1 {
			removeIsland(grid, nr, nc)
		}
	}
}

func isOnEndge(r, c, mr, mc int) bool {
	return r == mr || r == 0 || c == 0 || c == mc
}

func visitedMatrix(h, l int) [][]int {
	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, l)
	}
	return m
}

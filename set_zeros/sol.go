package set_zeros

func SetZeroes(matrix [][]int) {

	indices := [][]int{}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				indices = append(indices, []int{y, x})
			}
		}
	}

	for _, ii := range indices {
		for _, dir := range dirs {
			setZerosInDir(matrix, ii[0], ii[1], dir)
		}
	}

}

var dirs = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func setZerosInDir(matrix [][]int, y, x int, dir []int) {
	ny, nx := y+dir[0], x+dir[1]
	if ny < 0 || nx < 0 || ny >= len(matrix) || nx >= len(matrix[ny]) {
		return
	}
	matrix[ny][nx] = 0
	setZerosInDir(matrix, ny, nx, dir)

}

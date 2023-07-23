package search2dmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	if rows == 1 && cols == 1 {
		return matrix[0][0] == target
	}

	row := 0
	found := false
	for {
		rm := (row + rows) / 2

		if target < matrix[row][0] {
			return false
		}
		if target > matrix[rows-1][cols-1] {
			return false
		}
		if target >= matrix[rm][0] && target <= matrix[rm][cols-1] {
			found = true
			row = rm
			break
		}
		if row == rows {
			break
		}
		if target >= matrix[row][0] && target <= matrix[rm][cols-1] {
			rows = rm
		}
		if target > matrix[rm][cols-1] && target <= matrix[rows-1][cols-1] {
			row = rm
		}
	}
	if !found {
		return false
	}

	refArray := matrix[row]

	l, r := 0, cols

	for {
		m := (l + r) / 2
		if refArray[m] == target {
			return true
		}
		if l >= r {
			return false
		}
		if refArray[m] > target {
			r = m - 1
		}
		if refArray[m] < target {
			l = m + 1
		}
	}

}

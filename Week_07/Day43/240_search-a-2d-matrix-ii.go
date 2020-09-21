package Day43

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	i, j := m-1, 0
	for i >= 0 && j < n {
		leftBottom := matrix[i][j]
		if target == leftBottom {
			return true
		}
		if target < leftBottom {
			i--
		} else {
			j++
		}
	}
	return false
}

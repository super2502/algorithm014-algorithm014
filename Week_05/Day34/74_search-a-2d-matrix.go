package Day34

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	l := 0
	r := m * n - 1

	for l <= r {
		mid := l + ( r - l ) / 2
		midVal := matrix[mid/n][mid%n]
		if  midVal == target {
			return true
		} else if midVal < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
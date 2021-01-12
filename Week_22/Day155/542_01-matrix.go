package Day155

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	var bfs func(i, j int)
	bfs = func(i, j int) {
		if matrix[i][j] == 0 {
			return
		}
		level := 1
		for {
			found := false
			for k := 0; k < level; k++ {
				if i-k >= 0 && i-k < m && j-level+k >= 0 && j-level+k < n && matrix[i-k][j-level+k] == 0 {
					matrix[i][j] = level
					found = true
					break
				}
				if i-level+k >= 0 && i-level+k < m && j+k >= 0 && j+k < n && matrix[i-level+k][j+k] == 0 {
					matrix[i][j] = level
					found = true
					break
				}
				if i+k >= 0 && i+k < m && j+level-k >= 0 && j+level-k < n && matrix[i+k][j+level-k] == 0 {
					matrix[i][j] = level
					found = true
					break
				}
				if i+level-k >= 0 && i+level-k < m && j-k >= 0 && j-k < n && matrix[i+level-k][j-k] == 0 {
					matrix[i][j] = level
					found = true
					break
				}
			}
			if found {
				break
			}
			level++
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bfs(i, j)
		}
	}
	return matrix
}

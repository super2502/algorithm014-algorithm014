package Day154

func maxDistance(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}
	dictX, dictY := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	queue := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	level := 0
	for len(queue) > 0 {
		level++
		tmp := make([][]int, 0)
		for _, point := range queue {
			for d := 0; d < 4; d++ {
				nextX, nextY := point[0]+dictX[d], point[1]+dictY[d]
				if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || grid[nextX][nextY] != 0 {
					continue
				}
				grid[nextX][nextY] = level
				tmp = append(tmp, []int{nextX, nextY})
			}
		}
		queue = tmp
	}
	if level == 1 {
		return -1
	}
	return level - 1
}

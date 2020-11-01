package Day84

func uniquePathsIII(grid [][]int) int {
	var start []int
	length := 0
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				start = []int{i, j}
			} else if grid[i][j] == 0 {
				length++
			}
		}
	}
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	ret := 0
	var dfs func(path int, length int, x, y int)
	dfs = func(path int, length int, x, y int) {
		//fmt.Printf("try (%v)(%v), path %v / %v\n", x, y, path, length)
		if grid[x][y] == 2 {
			if path == length+1 {
				ret++
			}
			return
		}
		grid[x][y] = -1
		path++
		for d := 0; d < 4; d++ {
			nextX, nextY := x+dictX[d], y+dictY[d]
			if nextX < 0 || nextX == m || nextY < 0 || nextY == n || grid[nextX][nextY] == -1 {
				continue
			}
			dfs(path, length, nextX, nextY)
		}
		path--
		grid[x][y] = 0
	}
	dfs(0, length, start[0], start[1])
	return ret
}

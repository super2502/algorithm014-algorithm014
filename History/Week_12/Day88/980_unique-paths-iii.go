package Day88

func uniquePathsIII(grid [][]int) int {
	var start = make([]int, 2)
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	path := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				start[0], start[1] = i, j
			} else if grid[i][j] == 0 {
				path++
			}
		}
	}
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	ret := 0
	var dfs func(level int, i, j int)
	dfs = func(level int, i, j int) {
		//fmt.Printf("dfs called (%v)(%v)(%v)(%v)(%v)\n", level, i, j , grid[i][j], path)
		if level == path+1 {
			if grid[i][j] == 2 {
				ret++
			}
			return
		}
		if grid[i][j] != 0 {
			return
		}
		grid[i][j] = -1
		for d := 0; d < 4; d++ {
			nextI, nextJ := i+dictX[d], j+dictY[d]
			if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n {
				continue
			}
			dfs(level+1, nextI, nextJ)
		}
		grid[i][j] = 0
	}
	grid[start[0]][start[1]] = 0
	dfs(0, start[0], start[1])
	return ret
}

package Week_04

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs200(i, j, m, n, &grid)
			}
		}
	}
	return cnt
}

func dfs200(i, j, m, n int, grid *[][]byte) {
	(*grid)[i][j] = '0'
	if i > 0 && (*grid)[i-1][j] == '1' {
		dfs200(i-1, j, m, n, grid)
	}
	if i < m-1 && (*grid)[i+1][j] == '1' {
		dfs200(i+1, j, m, n, grid)
	}
	if j > 0 && (*grid)[i][j-1] == '1' {
		dfs200(i, j-1, m, n, grid)
	}
	if j < n-1 && (*grid)[i][j+1] == '1' {
		dfs200(i, j+1, m, n, grid)
	}
}

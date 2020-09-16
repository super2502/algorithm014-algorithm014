package Week_06

func minPathSum(grid [][]int) int {

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[0][0] = grid[0][0]
			for j := 1; j < n; j++ {
				dp[0][j] = dp[0][j-1] + grid[0][j]
			}
			continue
		}
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = min64(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func minPathSum1(grid [][]int) int {

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = -1
		if i == 0 {
			for j := 0; j <= n; j++ {
				dp[i][j] = -1
			}
			continue
		}
		for j := 1; j <= n; j++ {
			if dp[i-1][j] == -1 && dp[i][j-1] == -1 {
				dp[i][j] = grid[i-1][j-1]
			} else if dp[i-1][j] == -1 || dp[i][j-1] == -1 {
				dp[i][j] = max64(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			} else {
				dp[i][j] = min64(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func min64(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max64(a, b int) int {
	if a > b {
		return a
	}
	return b
}

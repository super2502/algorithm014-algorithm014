package Week_06

func uniquePaths(m int, n int) int {
	// i 行 j 列
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		if i == 0 {
			for j := 0; j < m; j++ {
				dp[i][j] = 1
			}
			continue
		}
		dp[i][0] = 1
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

package Day58

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for j := 1; j < m; j++ {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[n-1]
}

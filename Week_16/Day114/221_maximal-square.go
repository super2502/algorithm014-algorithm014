package Day114

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	maxEdge := 0
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				maxEdge = max(maxEdge, dp[i][j])
			}
		}
	}
	return maxEdge * maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

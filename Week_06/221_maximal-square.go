package Week_06

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	maxSide := 0
	for i := 0; i < m ; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min221(min221(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxSide = max221(dp[i][j], maxSide)
		}
	}
	return maxSide * maxSide
}

func max221(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min221(a, b int) int {
	if a < b {
		return a
	}
	return b
}
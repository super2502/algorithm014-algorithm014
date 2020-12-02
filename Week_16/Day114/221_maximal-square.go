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
	dp := make([]int, n+1)
	maxEdge := 0
	for i := 1; i <= m; i++ {
		leftUp := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = min(leftUp, min(dp[j], dp[j-1])) + 1
				maxEdge = max(maxEdge, dp[j])
			} else {
				dp[j] = 0
			}
			leftUp = tmp
		}
		//fmt.Printf("%+v\n", dp)
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

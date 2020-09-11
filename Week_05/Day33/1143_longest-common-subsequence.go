package Day33

func longestCommonSubsequence(text1 string, text2 string) int {

	m := len(text1)
	n := len(text2)
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	bs1 := []byte(text1)
	bs2 := []byte(text2)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if bs1[i-1] == bs2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

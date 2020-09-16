package Week_06

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max1143(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max1143(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

package Week_06

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m + 1)
	for i := 0; i <= m ; i++ {
		dp[i] = make([]int, n + 1)
		if i == 0 {
			for j := 0; j <= n ; j++ {
				dp[0][j] = j
			}
			continue
		}
		dp[i][0] = i
		for j := 1; j<=n ;j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min72(min72(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min72(a, b int) int {
	if a < b {
		return a
	}
	return b
}

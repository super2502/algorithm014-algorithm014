package Week_09

func numDistinct(s string, t string) int {
	m := len(t)
	n := len(s)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= m; i++ {
		for j := i; j <= n; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

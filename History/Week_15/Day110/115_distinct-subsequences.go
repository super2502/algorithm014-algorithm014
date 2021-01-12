package Day110

func numDistinct(s string, t string) int {
	m := len(t)
	n := len(s)
	if m == 0 {
		return 1
	}
	if n == 0 || m > n {
		return 0
	}
	dp := make([][]int, m+1)
	dp[0] = make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if t[i-1] == s[i-1] {
			dp[i][i] = dp[i-1][i-1]
		} else {
			dp[i][i] = 0
		}
		for j := i + 1; j <= n; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

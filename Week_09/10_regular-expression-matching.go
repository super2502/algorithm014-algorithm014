package Week_09

func isMatch10(s string, p string) bool {

	m := len(p)
	n := len(s)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		if p[i-1] == '*' {
			if i == 1 {
				return false
			}
			dp[i][0] = dp[i-2][0]
		}
		for j := 1; j <= n; j++ {
			if p[i-1] == '*' {
				if p[i-2] == '.' || p[i-2] == s[j-1] {
					dp[i][j] = dp[i][j-1] || dp[i-2][j-1] || dp[i-2][j]
				} else {
					dp[i][j] = dp[i-2][j]
				}
			} else {
				if p[i-1] == '.' || p[i-1] == s[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

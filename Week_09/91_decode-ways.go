package Week_09

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		if s[i-1] == '0' {
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			dp[i] = dp[i-2]
		} else {
			if (s[i-1] <= '6' && s[i-2] == '2') || (s[i-2] == '1') {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}

	return dp[n]
}

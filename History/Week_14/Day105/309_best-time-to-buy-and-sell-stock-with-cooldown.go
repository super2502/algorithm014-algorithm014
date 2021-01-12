package Day105

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	dp2 := 0
	for i := 1; i < n; i++ {
		dp2Tmp := dp2
		dp2 = dp1 + prices[i]
		dp1 = max(dp1, dp0-prices[i])
		dp0 = max(dp0, dp2Tmp)
	}
	return max(dp0, dp2)
}

func maxProfit1(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	dp[0] = []int{0, -prices[0], 0}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 3)
		dp[i][2] = dp[i-1][1] + prices[i]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
	}
	return max(dp[n-1][0], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

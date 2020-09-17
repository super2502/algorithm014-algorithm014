package Week_06

func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)
	dp[0][0] = 0
	dp[1][0] = -1 * prices[0]
	for i := 1; i < n; i++ {
		dp[0][i] = max714(dp[0][i-1], dp[1][i-1]+prices[i]-fee)
		dp[1][i] = max714(dp[1][i-1], dp[0][i-1]-prices[i])
	}
	return dp[0][n-1]
}

func max714(a, b int) int {
	if a > b {
		return a
	}
	return b
}

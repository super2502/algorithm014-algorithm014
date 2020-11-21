package Day104

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	ret := 0
	dp := make([][][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([][]int, n)
		for d := 0; d < n; d++ {
			dp[i][d] = make([]int, 2)
		}
	}
	dp[0][0][1] = -prices[0]
	for i := 1; i <= k; i++ {
		dp[i][0][1] = math.MinInt64
		for d := 1; d < n; d++ {
			dp[i][d][0] = max(dp[i][d-1][0], dp[i-1][d-1][1]+prices[d])
			dp[i-1][d][1] = max(dp[i-1][d-1][1], dp[i-1][d-1][0]-prices[d])
		}
		ret = max(ret, dp[i][n-1][0])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

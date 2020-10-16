package Day68

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp00 := 0
	dp10 := -1 * prices[0]
	dp01 := 0
	dp11 := math.MinInt64
	dp02 := math.MinInt64

	for i := 1; i < len(prices); i++ {
		dp02 = max(dp02, dp11+prices[i])
		dp11 = max(dp11, dp01-prices[i])
		dp01 = max(dp01, dp10+prices[i])
		dp10 = max(dp10, dp00-prices[i])
	}

	return max(max(dp00, dp01), dp02)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

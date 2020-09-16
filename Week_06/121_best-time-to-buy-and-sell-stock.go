package Week_06

import "math"

func maxProfit121(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	i00 := 0
	i01 := -1 * prices[0]
	i10 := math.MinInt64

	for i := 1; i < n; i++ {
		i10 = max121(i10, i01+prices[i])
		i01 = max121(i01, i00-prices[i])
	}
	return max121(i00, i10)
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}

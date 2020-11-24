package Day106

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	for i := 1; i < n; i++ {
		dp0Tmp := dp0
		dp0 = max(dp0, dp1+prices[i]-fee)
		dp1 = max(dp1, dp0Tmp-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

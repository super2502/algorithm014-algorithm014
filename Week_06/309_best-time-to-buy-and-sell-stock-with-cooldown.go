package Week_06

func maxProfit309(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp0, dp1, dp2 := 0, -1*prices[0], 0
	for i := 1; i < n; i++ {
		dp2Tmp := dp2
		dp2 = max309(dp0, dp1+prices[i])
		dp1 = max309(dp1, dp0-prices[i])
		dp0 = max309(dp0, dp2Tmp)
	}
	return max309(dp0, dp2)
}

func max309(a, b int) int {
	if a > b {
		return a
	}
	return b
}

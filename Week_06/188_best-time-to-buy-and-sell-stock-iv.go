package Week_06

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	if k > n/2 {
		sum := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				sum += prices[i] - prices[i-1]
			}
		}
		return sum
	}
	dp := make([][][]int, k+1)
	for j := 0; j <= k; j++ {
		dp[j] = make([][]int, 2)
		dp[j][0] = make([]int, n)
		dp[j][1] = make([]int, n)
		dp[j][1][0] = -1 * prices[0]
	}

	maxRet := 0
	for j := 0; j <= k; j++ {
		for i := 1; i < n; i++ {
			if j > 0 {
				dp[j][0][i] = max188(dp[j][0][i-1], dp[j-1][1][i-1]+prices[i])
			}
			dp[j][1][i] = max188(dp[j][1][i-1], dp[j][0][i-1]-prices[i])
		}
		maxRet = max188(maxRet, dp[j][0][n-1])
	}
	return maxRet
}

func max188(a, b int) int {
	if a > b {
		return a
	}
	return b
}

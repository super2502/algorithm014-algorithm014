package Day107

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin] == amount+1 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

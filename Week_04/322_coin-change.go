package Week_04

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for i := 1; i <= amount; i++ {
		//fmt.Printf("dp [%+v], i %v\n", dp, i)
		for _, coin := range coins {
			//fmt.Printf("test coin %v\n", coin)
			if i-coin < 0 {
				//fmt.Printf("drop coin %v\n", coin)
				continue
			}
			if dp[i-coin] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[i-coin] + 1
			} else {
				dp[i] = min322(dp[i], dp[i-coin]+1)
			}
		}
	}

	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}

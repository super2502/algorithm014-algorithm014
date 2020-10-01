package Week_08

func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}

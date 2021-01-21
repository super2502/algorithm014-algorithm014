package Day164

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j] * dp[i-j-1]
		}
		dp[i] = sum
	}
	//fmt.Printf("%+v\n", dp)
	return dp[n]
}

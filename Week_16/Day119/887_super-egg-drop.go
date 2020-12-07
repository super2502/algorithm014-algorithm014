package Day119

import "fmt"

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	dp[1] = make([]int, N+1)
	for j := 1; j <= N; j++ {
		dp[1][j] = j
	}
	fmt.Printf("%+v\n", dp[1])
	for i := 2; i <= K; i++ {
		dp[i] = make([]int, N+1)
		dp[i][1] = 1
		for j := 2; j <= N; j++ {
			dp[i][j] = 1 + dp[i-1][j-1] + dp[i][j-1]
		}
		fmt.Printf("%+v\n", dp[i])
	}
	for j := 1; j <= N; j++ {
		fmt.Printf("%v,%v,%v\n", j, dp[K][j], N)
		if dp[K][j] >= N {
			return j
		}
	}
	return -1
}

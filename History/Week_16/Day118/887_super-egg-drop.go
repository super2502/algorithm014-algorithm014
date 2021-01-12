package Day118

import "fmt"

func superEggDrop(K int, N int) int {

	dp := make([][]int, K+1)
	dp[1] = make([]int, N+1)
	for j := 1; j <= N; j++ {
		dp[1][j] = j
	}
	for i := 2; i <= K; i++ {
		dp[i] = make([]int, N+1)
		dp[i][1] = 1
		for j := 2; j <= N; j++ {
			l, r := 2, j-1
			for l <= r {
				mid := l + (r-l)>>1
				if dp[i][j-mid] == dp[i-1][mid-1] {
					l = mid
					break
				} else if dp[i][j-mid] < dp[i-1][mid-1] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			maxL := max(dp[i][j-l], dp[i-1][l-1])
			maxR := max(dp[i][j-r], dp[i-1][r-1])
			dp[i][j] = min(maxL, maxR) + 1
		}
		//for j := 2; j <= N; j++ {
		//	min := N+1
		//	for k := 1; k <= j-1; k++ {
		//		cnt := max(dp[i][j-k], dp[i-1][k-1])
		//		if min > cnt {
		//			min = cnt
		//		}
		//	}
		//	dp[i][j] = min + 1
		//}
	}
	for i := 0; i <= K; i++ {
		fmt.Printf("%+v\n", dp[i])
	}
	return dp[K][N]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

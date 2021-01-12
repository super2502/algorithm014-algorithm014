package Day149

import (
	"math"
)

func numSquares(n int) int {

	m := make(map[int]bool)
	x := 1
	for {
		y := x * x
		if y <= n {
			m[y] = true
		} else {
			break
		}
		x++
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if m[i] {
			dp[i] = 1
			continue
		}
		dp[i] = math.MaxInt64
		for j := 1; j <= i/2; j++ {
			dp[i] = min(dp[i], dp[j]+dp[i-j])
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

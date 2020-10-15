package Week_09

import "math"

func racecar(target int) int {

	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MaxInt64
		for forward := 1; (1<<forward)-1 < 2*i; forward++ {
			j := (1 << forward) - 1
			if j == i {
				dp[i] = forward
			} else if j > i {
				dp[i] = min(dp[i], forward+1+dp[j-i])
			} else {
				for back := 0; back < forward; back++ {
					k := (1 << back) - 1
					dp[i] = min(dp[i], forward+1+back+1+dp[i-j+k])
				}
			}
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

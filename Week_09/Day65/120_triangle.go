package Day65

import "math"

func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	minTotal := math.MaxInt64
	for _, sum := range dp[n-1] {
		minTotal = min(minTotal, sum)
	}
	return minTotal
}
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		leftUp := dp[0]
		dp[0] += triangle[i][0]
		for j := 1; j < i; j++ {
			tmp := dp[j]
			dp[j] = min(leftUp, dp[j]) + triangle[i][j]
			leftUp = tmp
		}
		dp[i] = leftUp + triangle[i][i]
	}
	minTotal := math.MaxInt64
	for _, sum := range dp {
		minTotal = min(minTotal, sum)
	}
	return minTotal
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package Day56

func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	dp := make([]int, n+1)
	maxLen := 0
	for i := 1; i <= m; i++ {
		leftUp := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if A[i-1] == B[j-1] {
				dp[j] = leftUp + 1
			} else {
				dp[j] = 0
			}
			maxLen = max(maxLen, dp[j])
			leftUp = tmp
		}
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

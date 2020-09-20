package Week_06

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			ret := 0
			if i == j || i == j+1 {
				continue
			}
			for k := i + 1; k < j; k++ {
				ret = max312(ret, nums[i]*nums[k]*nums[j]+dp[i][k]+dp[k][j])
			}
			dp[i][j] = ret
		}
	}
	return dp[0][n-1]
}

func max312(a, b int) int {
	if a > b {
		return a
	}
	return b
}

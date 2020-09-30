package Day52

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	first := nums[0]
	second := max198(nums[0], nums[1])
	for i := 2; i < n; i++ {
		cur := max198(first+nums[i], second)
		first = second
		second = cur
	}
	return second
}

func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max198(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max198(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max198(a, b int) int {
	if a > b {
		return a
	}
	return b
}

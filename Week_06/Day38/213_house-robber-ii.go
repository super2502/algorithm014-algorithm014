package Day38

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max213(rob0(nums[:len(nums)-1]), rob0(nums[1:len(nums)]))
}

func rob0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max213(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max213(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}

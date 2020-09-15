package Week_06

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxRet := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		maxRet = max53(maxRet, sum)
	}
	return maxRet
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}

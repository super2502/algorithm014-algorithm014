package Day46

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ret := math.MinInt64
	maxP := 1
	minP := 1
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			maxP, minP = minP, maxP
		}
		maxP = max(maxP*nums[i], nums[i])
		minP = min(minP*nums[i], nums[i])
		ret = max(ret, maxP)
	}
	return ret
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

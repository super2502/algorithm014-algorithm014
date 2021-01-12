package Day99

import "math"

func maxSubArray(nums []int) int {
	maxSub := math.MinInt64
	sum := 0
	for _, num := range nums {
		sum += num
		maxSub = max(maxSub, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSub
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

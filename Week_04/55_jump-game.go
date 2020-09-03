package Week_04

func canJump(nums []int) bool {
	maxDist := 0
	if len(nums) <= 1 {
		return true
	}

	for i := 0; i < len(nums); i++ {
		if maxDist < i {
			return false
		}
		maxDist = max(maxDist, i+nums[i])
		if maxDist >= len(nums)-1 {
			return true
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package Day122

func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	for i := 0; i < n; i++ {
		for nums[i] != i+1 && nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			idx := nums[i] - 1
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

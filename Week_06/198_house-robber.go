package Week_06

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		cur := max(first+nums[i], second)
		first = second
		second = cur
	}
	return second
}

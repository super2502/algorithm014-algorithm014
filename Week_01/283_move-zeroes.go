package Week_01

func moveZeroes(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	idx := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := idx; i < l; i++ {
		nums[i] = 0
	}
}

package Day14

func moveZeroes(nums []int) {
	idx := 0
	l := len(nums)
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

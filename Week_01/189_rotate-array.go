package Week_01

func rotate(nums []int, k int) {

	l := len(nums)
	if l == 0 {
		return
	}

	revert(nums, 0, l-1)
	revert(nums, 0, k%l-1)
	revert(nums, k%l, l-1)

}

func revert(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}

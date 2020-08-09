package Week_01

func moveZeroes(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	i := 0
	j := 0
	for j < l {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}

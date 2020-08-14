package Week_01

func removeDuplicates(nums []int) int {

	l := len(nums)
	if l <= 1 {
		return l
	}
	i, j := 0, 1
	for j < l {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}

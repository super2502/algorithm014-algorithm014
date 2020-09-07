package Day28

func majorityElement(nums []int) int {
	A := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			A = nums[i]
			count = 1
			continue
		}
		if nums[i] != A {
			count--
			continue
		}
		count++
	}
	return A
}

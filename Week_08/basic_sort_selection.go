package Week_08

func selectionSort(nums []int) []int {
	for i := 0 ; i < len(nums) ; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

package Week_08

func bubbleSort(nums []int) []int {
	for i:= 0; i< len(nums);i++ {
		for j:=1 ;j<len(nums) - i ;j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

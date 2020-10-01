package Week_08

func insertionSort(nums []int) {

	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i
		for j > 0 && tmp < nums[j-1] {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = tmp
	}
}

package Week_08

import (
	//"fmt"
)
func quickSort(nums []int, i, j int) {

	//fmt.Printf("in loop (%v)(%v)\n", i, j)
	if i >= j {
		return
	}
	pivot := partition(nums, i, j)
	//fmt.Printf("get pivot (%v)(%v)(%v) (%+v)\n", i, j, pivot, nums)
	quickSort(nums,i ,pivot-1)
	quickSort(nums, pivot+1, j)

}

func partition(nums []int, i, j int) int{

	pivot := j
	start := i
	for k := i; k < j; k++ {
		if nums[k] < nums[j] {
			nums[start], nums[k] = nums[k], nums[start]
			start++
		}
	}
	nums[pivot], nums[start] = nums[start], nums[pivot]
	return start
}
package Week_08

import (
//"fmt"
)

func mergeSort(nums []int, left, right int) {
	//fmt.Printf("loop in i(%v) j (%v) \n", i , j)
	if left >= right {
		return
	}

	mid := left + (right-left)>>1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {

	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	copy(nums[left:right+1], tmp)
}

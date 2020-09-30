package main

import (
	//"fmt"
)

func mergeSort(nums []int, i , j int) []int {
	//fmt.Printf("loop in i(%v) j (%v) \n", i , j)
	if i > j {
		return []int{}
	}
	if i == j {
		return []int{nums[i]}
	}
	mid := i + ( j - i ) >> 1
	left := mergeSort(nums, i, mid)
	right := mergeSort(nums, mid + 1 , j)

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	ret := make([]int, len(left) + len(right))

	i, j ,k := 0, 0 ,0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret[k] = left[i]
			i++
		} else {
			ret[k] = right[j]
			j++
		}
		k++
	}
	if  i == len(left) {
		for j < len(right) {
			ret[k] = right[j]
			j++
			k++
		}
	}
	if j == len(right) {
		for i < len(left) {
			ret[k] = left[i]
			i++
			k++
		}
	}

	return ret
}

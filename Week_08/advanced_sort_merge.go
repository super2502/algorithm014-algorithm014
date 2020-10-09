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

func mergeSortIteration(nums []int) {
	n := len(nums)
	step := 1
	for {
		//fmt.Printf("step: %v, before merge, %+v\n", step, nums)
		for i := 0; i < n ; i+=step<<1 {
			mergeIteration(nums, i, step)
		}
		//fmt.Printf("step: %v, after merge, %+v\n", step, nums)
		step = step << 1
		if step >= n {
			break
		}
	}

}

func mergeIteration(nums []int, start int, step int) {
	//fmt.Printf("step: %v, before merge, %+v\n", step, nums)
	//defer func() {
	//	fmt.Printf("step: %v, after merge, %+v\n", step, nums)
	//}()
	tmp := make([]int, 0, step * 2)
	if start + step >= len(nums) {
		return
	}

	end := start + 2 * step
	if end >= len(nums) {
		end = len(nums)
	}
	i, j := start, start + step
	for i < start + step && j < end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i < start + step {
		tmp = append(tmp, nums[i])
		i++
	}
	for j < end {
		tmp = append(tmp, nums[j])
		j++
	}
	copy(nums[start: end], tmp)
}

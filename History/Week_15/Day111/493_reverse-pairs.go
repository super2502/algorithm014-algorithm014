package Day111

import "fmt"

func reversePairs(nums []int) int {
	fmt.Printf("%+v\n", nums)
	ret := mergeSort(nums, 0, len(nums)-1)
	//fmt.Printf("%+v\n", nums)
	return ret
}

func mergeSort(nums []int, i, j int) int {
	if i >= j {
		return 0
	}
	mid := i + (j-i)>>1
	left := mergeSort(nums, i, mid)
	right := mergeSort(nums, mid+1, j)
	cnt := merge(nums, i, mid, j)
	return left + right + cnt
}
func merge(nums []int, i, mid, j int) int {
	//fmt.Printf("%+v, %v,%v,%v\n", nums,i ,mid, j)
	tmp := make([]int, j-i+1)
	first, second := i, mid+1
	cnt := 0
	for first <= mid && second <= j {
		if nums[first] > 2*nums[second] {
			cnt += mid - first + 1
			second++
		} else {
			first++
		}
	}
	first, second, k := i, mid+1, 0
	for first <= mid && second <= j {
		if nums[first] <= nums[second] {
			tmp[k] = nums[first]
			first++
		} else {
			tmp[k] = nums[second]
			second++
		}
		k++
	}
	for first <= mid {
		tmp[k] = nums[first]
		first++
		k++
	}
	for second <= j {
		tmp[k] = nums[second]
		second++
		k++
	}
	copy(nums[i:j+1], tmp)
	//fmt.Printf("%+v, %v\n", nums, cnt)
	return cnt
}

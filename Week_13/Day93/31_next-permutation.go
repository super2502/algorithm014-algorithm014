package Day93

import (
	"sort"
)

func nextPermutation(nums []int) {
	//fmt.Printf("%+v\n", nums)
	n := len(nums)
	if n == 0 {
		return
	}
	var stop int
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		} else {
			stop = i + 1
			break
		}
	}
	if stop == 0 {
		sort.Ints(nums)
		//fmt.Printf("%+v\n", nums)
		return
	}
	minHigher := stop
	for i := stop; i < n; i++ {
		if nums[i] > nums[stop-1] && nums[i] < nums[minHigher] {
			minHigher = i
		}
	}
	nums[stop-1], nums[minHigher] = nums[minHigher], nums[stop-1]
	sort.Ints(nums[stop:])
	//fmt.Printf("%+v\n", nums)
}

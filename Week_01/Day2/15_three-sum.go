package Day2

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	l := len(nums)
	if l < 3 {
		return ret
	}
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ret = append(ret, twoSum(nums[i], nums[i+1:l], 0)...)
	}
	return ret
}

func twoSum(head int, nums []int, sum int) [][]int {

	ret := make([][]int, 0)
	l := len(nums)
	if l < 2 {
		return ret
	}
	c := sum - head
	for i, j := 0, l-1; i < j; {
		if nums[i]+nums[j] == c {
			ret = append(ret, []int{head, nums[i], nums[j]})
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			i++
			j--
		}

		if nums[i]+nums[j] > c {
			j--
		} else if nums[i]+nums[j] < c {
			i++
		}
	}
	return ret

}

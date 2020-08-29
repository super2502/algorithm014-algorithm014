package Day20

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)

	l := len(nums)
	if l < 3 {
		return ret
	}
	for i := 0; i < l; {
		tar := 0 - nums[i]
		j := i + 1
		k := l - 1
		for j < k {
			sum := nums[j] + nums[k]
			if sum == tar {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum < tar {
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else {
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			}
		}
		for i < l-1 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}

	return ret
}

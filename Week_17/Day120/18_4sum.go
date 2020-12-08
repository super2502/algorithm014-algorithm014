package Day120

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	//fmt.Printf("%+v\n", nums)
	for i := 0; i <= len(nums)-3; i++ {
		threeSum(nums[i+1:], target, nums[i], &ret)
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		//fmt.Printf("after (%v) get (%+v)\n", i, ret)
	}
	//fmt.Printf("%+v\n",ret)
	return ret
}
func threeSum(nums []int, target int, first int, ret *[][]int) {
	//fmt.Printf("three sum get (%+v)(%v)(%v)\n", nums, target ,first)
	for i := 0; i < len(nums)-2; i++ {
		twoSum(nums[i+1:], target, first, nums[i], ret)
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
}
func twoSum(nums []int, target int, first, second int, ret *[][]int) {
	//fmt.Printf("two sum get (%+v) (%v)(%v)(%v)\n", nums, target ,first, second)
	i, j := 0, len(nums)-1
	for i < j {
		sum := first + second + nums[i] + nums[j]
		if sum == target {
			//fmt.Printf("two sum hit (%v)(%v)(%v)(%v)\n", first,second,nums[i],nums[j])
			*ret = append(*ret, []int{first, second, nums[i], nums[j]})
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
			for j > 0 && nums[j] == nums[j-1] {
				j--
			}
			i++
			j--
		} else if sum < target {
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
			i++
		} else {
			for j > 0 && nums[j] == nums[j-1] {
				j--
			}
			j--
		}
	}
}

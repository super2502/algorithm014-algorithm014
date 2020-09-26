package Day48

import (
	"sort"
	"math"
	//"fmt"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) < 3 {
		return 0
	}
	ret := math.MaxInt64
	sum := 0
	for i := 0; i < len(nums); i++ {
		targetTwo := target - nums[i]
		j := i+1
		k := len(nums) - 1
		for j < k {
			//fmt.Printf("check %v, %v, %v (%v)\n", nums[i], nums[j], nums[k], targetTwo)
			two := nums[j] + nums[k]
			if two == targetTwo {
				return target
			}
			if two < targetTwo {
				if targetTwo - two < ret {
					ret = targetTwo - two
					sum = nums[i] + nums[j] + nums[k]
				}
				//ret = min(ret, targetTwo - two)
				for j < k && nums[j] == nums[j+1]{
					j++
				}
				j++
			} else {
				if two - targetTwo < ret {
					ret = two - targetTwo
					sum = nums[i] + nums[j] + nums[k]
				}
				//ret = min(ret, two - targetTwo)
				for j < k && nums[k] == nums[k-1]{
					k--
				}
				k--
			}
		}
		for i < len(nums) - 2 && nums[i] == nums[i+1] {
			i++
		}
	}


	return sum

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
//
//
//示例：
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//提示：
//
//3 <= nums.length <= 10^3
//-10^3 <= nums[i] <= 10^3
//-10^4 <= target <= 10^4

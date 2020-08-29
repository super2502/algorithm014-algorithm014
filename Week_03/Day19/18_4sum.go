package Day19

import "sort"

var ret [][]int
var gnums []int

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	gnums = nums
	ret = make([][]int, 0)
	for i := 0; i < len(gnums); {
		threeSum(i, target, &ret)
		for i < len(gnums)-1 && gnums[i] == gnums[i+1] {
			i++
		}
		i++
	}
	return ret
}

func threeSum(start int, target int, ret *[][]int) {
	if start > len(gnums)-3 {
		return
	}
	four := gnums[start]
	for i := start + 1; i < len(gnums); {
		twoSum(four, i, target, ret)
		for i < len(gnums)-1 && gnums[i] == gnums[i+1] {
			i++
		}
		i++
	}
}

func twoSum(four int, start int, target int, ret *[][]int) {

	sum := target - four - gnums[start]
	if start > len(gnums)-2 {
		return
	}

	i, j := start+1, len(gnums)-1
	for i < j {
		s := gnums[i] + gnums[j]
		if s == sum {
			*ret = append(*ret, append([]int{four, gnums[start]}, gnums[i], gnums[j]))
			for i < j && gnums[i] == gnums[i+1] {
				i++
			}
			for i < j && gnums[j] == gnums[j-1] {
				j--
			}
			i++
			j--
		} else if s < sum {
			for i < j && gnums[i] == gnums[i+1] {
				i++
			}
			i++
		} else {
			for i < j && gnums[j] == gnums[j-1] {
				j--
			}
			j--
		}
	}
}

// func sumArr(nums []int) int {
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     return sum
// }

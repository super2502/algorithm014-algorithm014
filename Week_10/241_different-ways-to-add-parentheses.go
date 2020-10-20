package Week_10

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	nums := parse(input)
	//for _, num := range nums {
	//	fmt.Printf("%+v\n", num)
	//}
	return helper(nums)
}

func helper(nums []*node) []int {
	if len(nums) == 1 {
		return []int{nums[0].num}
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		lefts := helper(nums[:i+1])
		rights := helper(nums[i+1:])
		symbol := nums[i].symbol
		for _, left := range lefts {
			for _, right := range rights {
				switch symbol {
				case '+':
					ret = append(ret, left+right)
				case '-':
					ret = append(ret, left-right)
				case '*':
					ret = append(ret, left*right)
				}
			}
		}
	}
	return ret
}

type node struct {
	num    int
	symbol byte
}

func parse(input string) []*node {
	nums := make([]*node, 0)

	num := ""
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			num += string(input[i])
		} else {
			n, _ := strconv.Atoi(num)
			nums = append(nums, &node{
				num:    n,
				symbol: input[i],
			})
			num = ""
		}
	}
	n, _ := strconv.Atoi(num)
	nums = append(nums, &node{
		num: n,
	})

	return nums
}

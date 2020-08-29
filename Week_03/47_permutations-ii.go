package Week_03

import (
	"fmt"
	"sort"
)

var ret47 [][]int

func permuteUnique(nums []int) [][]int {
	ret47 = make([][]int, 0)
	sort.Ints(nums)
	fmt.Printf("nums %v\n", nums)
	dfs47([]int{}, nums, len(nums))
	return ret47
}

func dfs47(path []int, nums []int, l int) {
	fmt.Printf("called with path (%+v), nums(%+v)\n", path, nums)
	if len(path) == l {
		p := make([]int, len(path))
		copy(p, path)
		ret47 = append(ret47, p)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs47(path, removeI(nums, i), l)
		path = path[:len(path)-1]
	}
}

func removeI(nums []int, i int) []int {
	first := make([]int, i)
	copy(first, nums[:i])

	return append(first, nums[i+1:]...)
}

package Week_03

import (
	"fmt"
	"sort"
)

// 这里用的方法是抠掉当前用的数字，然后将剩余部分向下递归，规避了visited乱套的问题
// 如果使用visited，注意下面用法
//if (*visited)[i] {
//	continue
//}
//if i > 0 && nums[i] == nums[i - 1] && !(*visited)[i-1] {
//	continue
//}
// 这里!(*visited)[i-1] 是因为只有回溯的兄弟节点才会把自己的visited去掉，这样证明这个节点就是兄弟节点
// 而兄弟节点值相同的正是要剪枝剪掉的，其余还在visited里的节点是父+节点，当前节点不能被剪掉还得继续拿去递归
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

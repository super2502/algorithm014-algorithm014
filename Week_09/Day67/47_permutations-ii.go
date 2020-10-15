package Day67

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	visited := make([]bool, len(nums))
	dfs([]int{}, nums, &ret, visited)
	return ret
}

func dfs(path []int, nums []int, ret *[][]int, visited []bool) {
	if len(path) == len(nums) {
		p := make([]int, len(path))
		copy(p, path)
		*ret = append(*ret, p)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		path = append(path, nums[i])
		visited[i] = true
		dfs(path, nums, ret, visited)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

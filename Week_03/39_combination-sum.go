package Week_03

// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。

var ret39 [][]int

func combinationSum(candidates []int, target int) [][]int {
	ret39 = make([][]int, 0)
	dfs39([]int{}, 0, candidates, target)
	return ret39
}

func dfs39(path []int, start int, candidates []int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		p := make([]int, len(path))
		copy(p, path)
		ret39 = append(ret39, p)
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs39(path, start, candidates[i:], target-candidates[i])
		path = path[:len(path)-1]
	}
}

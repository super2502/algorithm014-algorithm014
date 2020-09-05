package Week_03

// 细节是魔鬼： i > start
import "sort"

var ret40 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	ret40 = make([][]int, 0)
	sort.Ints(candidates)
	dfs40([]int{}, 0, candidates, target)
	return ret40
}

func dfs40(path []int, start int, candidates []int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		p := make([]int, len(path))
		copy(p, path)
		ret40 = append(ret40, p)
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		dfs40(path, i+1, candidates, target-candidates[i])
		path = path[:len(path)-1]
	}

}

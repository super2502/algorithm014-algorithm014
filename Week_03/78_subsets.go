package Week_03

var ret78 [][]int

func subsets(nums []int) [][]int {
	ret78 = make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ret78
	}
	dfs78([]int{}, nums, l, 0)

	return ret78
}

func dfs78(path []int, nums []int, l int, level int) {
	if level == l {
		p := make([]int, len(path))
		copy(p, path)
		ret78 = append(ret78, p)
		return
	}

	num := nums[level]
	dfs78(path, nums, l, level+1)
	path = append(path, num)
	dfs78(path, nums, l, level+1)
}

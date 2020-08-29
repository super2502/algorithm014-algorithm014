package Week_03

var ret46 [][]int

func permute(nums []int) [][]int {
	ret46 = make([][]int, 0)
	visited := make([]bool, len(nums))
	dfs46([]int{}, nums, visited)
	return ret46
}

func dfs46(path []int, nums []int, visited []bool) {
	if len(path) == len(nums) {
		p := make([]int, len(path))
		copy(p, path)
		ret46 = append(ret46, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		path = append(path, nums[i])
		visited[i] = true
		dfs46(path, nums, visited)
		visited[i] = false
		path = path[:len(path)-1]
	}
}

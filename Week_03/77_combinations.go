package Week_03

var ret77 [][]int

func combine(n int, k int) [][]int {
	ret77 = make([][]int, 0)
	dfs77([]int{}, 1, n, k)
	return ret77
}

func dfs77(path []int, start int, n int, k int) {
	if len(path) == k {
		p := make([]int, k)
		copy(p, path)
		ret77 = append(ret77, p)
		return
	}
	for i := start; i <= n-k+len(path)+1; i++ {
		path = append(path, i)
		dfs77(path, i+1, n, k)
		path = path[:len(path)-1]
	}
}

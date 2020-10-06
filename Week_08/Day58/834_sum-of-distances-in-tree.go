package Day58

func sumOfDistancesInTree(N int, edges [][]int) []int {
	ret := make([]int, N)
	aMatrix := make([][]int, N) // 邻接矩阵超内存，这实际是个邻接表
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if len(aMatrix[x]) == 0 {
			aMatrix[x] = make([]int, 0)
		}
		aMatrix[x] = append(aMatrix[x], y)
		if len(aMatrix[y]) == 0 {
			aMatrix[y] = make([]int, 0)
		}
		aMatrix[y] = append(aMatrix[y], x)
	}
	depth := make([]int, N)
	nodes := make([]int, N)
	var dfs func(i int, iParent int)
	dfs = func(i int, iParent int) {
		nodes[i] = 1
		children := aMatrix[i]
		for _, child := range children {
			if child == iParent {
				continue
			}
			depth[child] = depth[i] + 1
			dfs(child, i)
			nodes[i] += nodes[child]
		}
	}
	dfs(0, -1)
	ret[0] = sum(depth)
	var dfs2 func(i, iParent int)
	dfs2 = func(i, iParent int) {
		children := aMatrix[i]
		for _, child := range children {
			if child == iParent {
				continue
			}
			ret[child] = ret[i] + N - 2*nodes[child]
			dfs2(child, i)
		}
	}
	dfs2(0, -1)
	return ret
}
func sum(nums []int) int {
	a := 0
	for _, num := range nums {
		a += num
	}
	return a
}

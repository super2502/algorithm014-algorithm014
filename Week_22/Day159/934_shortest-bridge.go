package Day159

func shortestBridge(A [][]int) int {
	n := len(A)
	queue := make([][]int, 0)
	dictX, dictY := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		queue = append(queue, []int{i, j})
		A[i][j] = 2
		for d := 0; d < 4; d++ {
			nextX, nextY := i+dictX[d], j+dictY[d]
			if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || A[nextX][nextY] != 1 {
				continue
			}
			dfs(nextX, nextY)
		}
	}
	stop := false
	for i := 0; i < n; i++ {
		if stop {
			break
		}
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				dfs(i, j)
				stop = true
				break
			}
		}
	}
	//for i := 0; i < n ; i++ {
	//	fmt.Printf("%+v\n", A[i])
	//}
	level := 0
	for len(queue) > 0 {
		tmp := make([][]int, 0)
		for _, point := range queue {
			i, j := point[0], point[1]
			for d := 0; d < 4; d++ {
				nextX, nextY := i+dictX[d], j+dictY[d]
				if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || A[nextX][nextY] == 2 {
					continue
				}
				if A[nextX][nextY] == 1 {
					return level
				}
				A[nextX][nextY] = 2
				tmp = append(tmp, []int{nextX, nextY})
			}
		}
		queue = tmp
		level++
	}
	return level
}

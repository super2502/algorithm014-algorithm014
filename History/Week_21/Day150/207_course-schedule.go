package Day150

func canFinish(numCourses int, prerequisites [][]int) bool {
	cMap := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		cMap[i] = make(map[int]bool)
	}
	for _, pre := range prerequisites {
		cMap[pre[1]][pre[0]] = true
	}
	visited := make([]bool, numCourses)
	var dfs func(i int, onStack []bool) bool
	dfs = func(i int, onStack []bool) bool {
		if len(cMap[i]) == 0 {
			return true
		}
		visited[i] = true
		for next := range cMap[i] {
			if onStack[next] {
				return false
			}
			if visited[next] {
				continue
			}
			visited[next] = true
			onStack[next] = true
			if !dfs(next, onStack) {
				return false
			}
			onStack[next] = false
		}
		return true
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] {
			continue
		}
		onStack := make([]bool, numCourses)
		onStack[i] = true
		if !dfs(i, onStack) {
			return false
		}
	}
	return true
}

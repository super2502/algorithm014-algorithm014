package Day107

import "fmt"

func canFinish(jobs int, prerequisites [][]int) bool {
	nexts := make(map[int]map[int]struct{})
	for _, pair := range prerequisites {
		if len(nexts[pair[1]]) == 0 {
			nexts[pair[1]] = make(map[int]struct{})
		}
		nexts[pair[1]][pair[0]] = struct{}{}
	}
	fmt.Printf("nexts (%+v)\n", nexts)
	for j := range nexts {
		visited := make(map[int]bool)
		onStack := make(map[int]bool)
		ok := dfs(j, visited, onStack, nexts)
		if !ok {
			return false
		}
	}
	return true
}

func dfs(j int, visited, onStack map[int]bool, nexts map[int]map[int]struct{}) bool {
	fmt.Printf("dfs (%v)(%+v)(%+v)\n", j, visited, onStack)

	visited[j] = true
	onStack[j] = true
	for nextJ := range nexts[j] {
		if onStack[nextJ] {
			return false
		}
		if visited[nextJ] {
			continue
		}
		ok := dfs(nextJ, visited, onStack, nexts)
		if !ok {
			return false
		}
	}
	onStack[j] = false
	return true
}

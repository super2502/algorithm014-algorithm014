package Day150

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	line := make(map[string]bool)
	//for _, query := range queries {
	//	line[query[0]] = struct{}{}
	//	line[query[1]] = struct{}{}
	//}
	for _, equation := range equations {
		line[equation[0]] = true
		line[equation[1]] = true
	}
	n := len(line)
	lineSlice := make([]string, 0, n)
	for s := range line {
		lineSlice = append(lineSlice, s)
	}
	lineKeyMap := make(map[string]int)
	for idx, s := range lineSlice {
		lineKeyMap[s] = idx
	}
	arr := make([][]*node, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]*node, n)
		arr[i][i] = &node{num: 1.0}
	}
	for idx, equations := range equations {
		arr[lineKeyMap[equations[0]]][lineKeyMap[equations[1]]] = &node{num: values[idx]}
		arr[lineKeyMap[equations[1]]][lineKeyMap[equations[0]]] = &node{num: 1 / values[idx]}
	}
	var dfs func(start, end int, visited []bool) (bool, float64)
	dfs = func(start, end int, visited []bool) (bool, float64) {
		if arr[start][end] != nil {
			return true, arr[start][end].num
		}
		if visited[end] {
			return false, -1.0
		}
		for i := 0; i < n; i++ {
			if arr[start][i] == nil || visited[i] {
				continue
			}
			visited[i] = true
			ok, prod := dfs(i, end, visited)
			if ok {
				return true, arr[start][i].num * prod
			}
		}
		return false, -1.0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited := make([]bool, n)
			visited[i] = true
			ok, prod := dfs(i, j, visited)
			if ok {
				arr[i][j] = &node{num: prod}
			}
		}
	}
	ret := make([]float64, len(queries))
	for idx, query := range queries {
		if !line[query[0]] || !line[query[1]] {
			ret[idx] = -1.0
		} else if arr[lineKeyMap[query[0]]][lineKeyMap[query[1]]] == nil {
			ret[idx] = -1.0
		} else {
			ret[idx] = arr[lineKeyMap[query[0]]][lineKeyMap[query[1]]].num
		}
	}
	return ret
}

type node struct {
	num     float64
	visited bool
}

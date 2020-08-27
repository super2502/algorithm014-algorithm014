package Week_03

var ret22 []string

func generateParenthesis(n int) []string {
	ret22 = make([]string, 0)
	dfs22([]byte{}, n, 0, 0)

	return ret22
}

func dfs22(path []byte, n int, left, right int) {
	if len(path) == 2*n {
		ret22 = append(ret22, string(path))
		return
	}
	if left < n {
		path = append(path, '(')
		dfs22(path, n, left+1, right)
		path = path[:len(path)-1]
	}
	if right < left {
		path = append(path, ')')
		dfs22(path, n, left, right+1)
		path = path[:len(path)-1]
	}
}

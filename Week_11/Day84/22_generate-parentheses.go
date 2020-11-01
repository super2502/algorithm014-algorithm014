package Day84

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	var dfs func(path []byte, n int, left, right int)
	dfs = func(path []byte, n int, left, right int) {
		if right == n {
			ret = append(ret, string(path))
			return
		}
		if left < n {
			path = append(path, '(')
			dfs(path, n, left+1, right)
			path = path[:len(path)-1]
		}
		if right < left {
			path = append(path, ')')
			dfs(path, n, left, right+1)
		}
	}
	dfs([]byte{}, n, 0, 0)
	return ret
}

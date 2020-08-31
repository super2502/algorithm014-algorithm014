package Week_04

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	dfs22([]byte{}, n, 0, 0, &ret)
	return ret
}

func dfs22(path []byte, n int, left, right int, ret *[]string) {
	if left == n && right == n {
		*ret = append(*ret, string(path))
		return
	}
	if left < n {
		path = append(path, '(')
		dfs22(path, n, left+1, right, ret)
		path = path[:len(path)-1]
	}
	if right < left {
		path = append(path, ')')
		dfs22(path, n, left, right+1, ret)
	}
}

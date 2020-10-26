package Day78

func findWords(board [][]byte, words []string) []string {
	ret := make([]string, 0)
	retMap := make(map[string]struct{})
	m := len(board)
	if m == 0 {
		return ret
	}
	n := len(board[0])
	if n == 0 {
		return ret
	}
	t := &trie{}
	for _, word := range words {
		t.insert(word)
	}
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	var dfs func(path []byte, i, j int)
	dfs = func(path []byte, i, j int) {
		tmp := board[i][j]
		path = append(path, tmp)
		str := string(path)
		if !t.prefix(str) {
			return
		}
		if t.match(string(path)) {
			retMap[str] = struct{}{}
		}
		board[i][j] = '.'
		for d := 0; d < 4; d++ {
			nextI, nextJ := i+dictX[d], j+dictY[d]
			if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || board[nextI][nextJ] == '.' {
				continue
			}
			dfs(path, nextI, nextJ)
		}
		board[i][j] = tmp
		path = path[:len(path)-1]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs([]byte{}, i, j)
		}
	}
	for str := range retMap {
		ret = append(ret, str)
	}
	return ret
}

type trie struct {
	isWord   bool
	children [26]*trie
}

func (t *trie) insert(word string) {
	root := t
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			root.children[key] = &trie{}
		}
		root = root.children[key]
	}
	root.isWord = true
}
func (t *trie) match(word string) bool {
	root := t
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return root.isWord
}
func (t *trie) prefix(word string) bool {
	root := t
	for i := 0; i < len(word); i++ {
		key := int(word[i] - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return true
}

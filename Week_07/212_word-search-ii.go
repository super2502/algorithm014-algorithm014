package Week_07

// 这个问题更优解貌似是把trie树的查找集成到dfs里，这样不用每次从root开始查，不过回溯就有点复杂了，还得把trie也回溯上去？

func findWords(board [][]byte, words []string) []string {
	ret := make([]string, 0)
	m := len(board)
	if m == 0 {
		return ret
	}
	n := len(board[0])
	if n == 0 {
		return ret
	}

	tr := &trie{}
	for _, word := range words {
		tr.insert(word)
	}
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	retMap := make(map[string]struct{})

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs212([]byte{}, &board, retMap, tr, dictX, dictY, i, j, m, n)
		}
	}
	for w := range retMap {
		ret = append(ret, w)
	}
	return ret
}

func dfs212(path []byte, board *[][]byte, retMap map[string]struct{}, tr *trie, dictX, dictY []int, x, y, m, n int) {
	tmp := (*board)[x][y]
	path = append(path, tmp)
	if !tr.beginWith(string(path)) { // 剪枝
		path = path[:len(path)-1]
		return
	}
	if tr.search(string(path)) {
		retMap[string(path)] = struct{}{}
	}
	(*board)[x][y] = 'X'
	for d := 0; d < 4; d++ {
		nextX, nextY := x+dictX[d], y+dictY[d]
		if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || (*board)[nextX][nextY] == 'X' {
			continue
		}
		dfs212(path, board, retMap, tr, dictX, dictY, nextX, nextY, m, n)

	}
	path = path[:len(path)-1] // 回溯
	(*board)[x][y] = tmp
}

type trie struct {
	isEnd    bool
	children [26]*trie
}

func (tr *trie) insert(word string) {
	root := tr
	for _, b := range []byte(word) {
		key := int(b - 'a')
		if root.children[key] == nil {
			root.children[key] = &trie{}
		}
		root = root.children[key]
	}
	root.isEnd = true
}

func (tr *trie) search(word string) bool {
	root := tr
	for _, b := range []byte(word) {
		key := int(b - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return root.isEnd
}

func (tr *trie) beginWith(prefix string) bool {
	root := tr
	for _, b := range []byte(prefix) {
		key := int(b - 'a')
		if root.children[key] == nil {
			return false
		}
		root = root.children[key]
	}
	return true
}

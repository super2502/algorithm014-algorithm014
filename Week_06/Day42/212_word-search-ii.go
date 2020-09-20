package Day42

func findWords(board [][]byte, words []string) []string {
	ret := make(map[string]struct{})
	dictX := []int{0, 1, 0, -1}
	dictY := []int{1, 0, -1, 0}
	trie := &Trie{
		children: make(map[byte]*Trie),
	}
	for _, word := range words {
		trie.insert(word)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs([]byte{}, board, i, j, &dictX, &dictY, trie, ret)
		}
	}
	ret1 := make([]string, 0, len(ret))
	for s := range ret {
		ret1 = append(ret1, s)
	}
	return ret1
}

func dfs(path []byte, board [][]byte, i, j int, dictX, dictY *[]int, trie *Trie, ret map[string]struct{}) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' {
		return
	}
	tmp := board[i][j]
	path = append(path, tmp)
	s := string(path)
	if !trie.prefix(s) {
		return
	}
	if trie.match(s) {
		ret[s] = struct{}{}
	}
	board[i][j] = '#'
	for k := 0; k < 4; k++ {
		dfs(path, board, i+(*dictX)[k], j+(*dictY)[k], dictX, dictY, trie, ret)
	}
	board[i][j] = tmp
}

type Trie struct {
	isEnd    bool
	children map[byte]*Trie
}

func (t *Trie) insert(s string) {
	root := t
	for _, b := range []byte(s) {
		if _, ok := root.children[b]; !ok {
			child := &Trie{
				children: make(map[byte]*Trie),
			}
			root.children[b] = child
		}
		root = root.children[b]
	}
	root.isEnd = true
}

func (t *Trie) match(s string) bool {
	root := t
	for _, b := range []byte(s) {
		if _, ok := root.children[b]; !ok {
			return false
		} else {
			root = root.children[b]
		}
	}
	return root.isEnd
}

func (t *Trie) prefix(s string) bool {
	root := t
	for _, b := range []byte(s) {
		if _, ok := root.children[b]; !ok {
			return false
		} else {
			root = root.children[b]
		}
	}
	return true
}

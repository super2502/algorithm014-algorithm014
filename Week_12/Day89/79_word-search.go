package Day89

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	dictI := []int{-1, 0, 1, 0}
	dictJ := []int{0, 1, 0, -1}
	var dfs func(level int, i, j int) bool
	dfs = func(level int, i, j int) bool {
		if board[i][j] != word[level] {
			return false
		}
		if level == len(word)-1 {
			return true
		}
		char := board[i][j]
		board[i][j] = '.'
		for d := 0; d < 4; d++ {
			nextI, nextJ := i+dictI[d], j+dictJ[d]
			if nextI < 0 || nextI == m || nextJ < 0 || nextJ == n || board[nextI][nextJ] == '.' {
				continue
			}
			if dfs(level+1, nextI, nextJ) {
				return true
			}
		}
		board[i][j] = char

		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}

package Week_07

func solveSudoku(board [][]byte) {
	rowMaps, colMaps, boxMaps := initMaps(board)
	dfs(0, board, rowMaps, colMaps, boxMaps)
}

func dfs(level int, board [][]byte, rowMaps, colMaps, boxMaps []map[byte]struct{}) bool {
	if level == 81 {
		return true
	}
	i, j := level/9, level%9
	boxIdx := i/3*3 + j/3
	if board[i][j] != '.' {
		return dfs(level+1, board, rowMaps, colMaps, boxMaps)
	}
	for k := 1; k <= 9; k++ {
		b := byte(k + '0')
		if !validate(rowMaps, colMaps, boxMaps, i, j, b) {
			continue
		}
		rowMaps[i][b] = struct{}{}
		colMaps[j][b] = struct{}{}
		boxMaps[boxIdx][b] = struct{}{}
		board[i][j] = b
		if dfs(level+1, board, rowMaps, colMaps, boxMaps) {
			return true
		}
		board[i][j] = '.'
		delete(rowMaps[i], b)
		delete(colMaps[j], b)
		delete(boxMaps[boxIdx], b)
	}
	return false
}

func validate(rowMaps, colMaps, boxMaps []map[byte]struct{}, i, j int, b byte) bool {
	if _, ok := rowMaps[i][b]; ok {
		return false
	}
	if _, ok := colMaps[j][b]; ok {
		return false
	}
	boxIdx := i/3*3 + j/3
	if _, ok := boxMaps[boxIdx][b]; ok {
		return false
	}
	return true
}
func initMaps(board [][]byte) ([]map[byte]struct{}, []map[byte]struct{}, []map[byte]struct{}) {
	rowMaps := make([]map[byte]struct{}, 9)
	colMaps := make([]map[byte]struct{}, 9)
	boxMaps := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		rowMaps[i] = make(map[byte]struct{})
		colMaps[i] = make(map[byte]struct{})
		boxMaps[i] = make(map[byte]struct{})
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			rowMaps[i][board[i][j]] = struct{}{}
			colMaps[j][board[i][j]] = struct{}{}
			boxIdx := i/3*3 + j/3
			boxMaps[boxIdx][board[i][j]] = struct{}{}
		}
	}
	return rowMaps, colMaps, boxMaps
}

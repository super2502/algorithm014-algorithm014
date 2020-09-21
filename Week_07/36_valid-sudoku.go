package Week_07

func isValidSudoku(board [][]byte) bool {
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
			if _, ok := rowMaps[i][board[i][j]]; ok {
				return false
			}
			rowMaps[i][board[i][j]] = struct{}{}
			if _, ok := colMaps[j][board[i][j]]; ok {
				return false
			}
			colMaps[j][board[i][j]] = struct{}{}
			boxIdx := i/3*3 + j/3
			if _, ok := boxMaps[boxIdx][board[i][j]]; ok {
				return false
			}
			boxMaps[boxIdx][board[i][j]] = struct{}{}
		}
	}
	return true
}

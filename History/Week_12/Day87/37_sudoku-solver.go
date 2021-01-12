package Day87

func solveSudoku(board [][]byte) {
	rowMap := [9][9]bool{}
	colMap := [9][9]bool{}
	boxMap := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '1')
			rowMap[i][num] = true
			colMap[j][num] = true
			boxIdx := i/3*3 + j/3
			boxMap[boxIdx][num] = true
		}
	}

	var dfs func(level int) bool
	dfs = func(level int) bool {
		if level == 81 {
			return true
		}
		i, j := level/9, level%9
		if board[i][j] != '.' {
			return dfs(level + 1)
		}
		boxIdx := i/3*3 + j/3
		for k := 0; k < 9; k++ {
			if rowMap[i][k] || colMap[j][k] || boxMap[boxIdx][k] {
				continue
			}
			rowMap[i][k] = true
			colMap[j][k] = true
			boxMap[boxIdx][k] = true
			board[i][j] = byte(k + '1')
			if dfs(level + 1) {
				return true
			}
			rowMap[i][k] = false
			colMap[j][k] = false
			boxMap[boxIdx][k] = false
			board[i][j] = '.'
		}
		return false
	}
	dfs(0)
	//
	//for i:= 0; i< 9;i++ {
	//	fmt.Printf("%+v\n", string(board[i]))
	//}
}

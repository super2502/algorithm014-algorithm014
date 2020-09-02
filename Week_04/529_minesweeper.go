package Week_04

func updateBoard(board [][]byte, click []int) [][]byte {
	m := len(board)
	if m == 0 {
		return board
	}
	n := len(board[0])
	if n == 0 {
		return board
	}
	dirX := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dirY := []int{1, 1, 1, 0, 0, -1, -1, -1}

	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	dfs(x, y, &board, m, n, &dirX, &dirY)
	return board
}

func dfs(x, y int, board *[][]byte, m, n int, dirX, dirY *[]int) {
	cnt := 0
	(*board)[x][y] = '#'
	for i := 0; i < 8; i++ {
		px := x + (*dirX)[i]
		py := y + (*dirY)[i]
		if px < 0 || px >= m || py < 0 || py >= n {
			continue
		}
		if (*board)[px][py] == 'M' {
			cnt++
		}
	}
	if cnt != 0 {
		(*board)[x][y] = byte(cnt + '0')
		return
	}

	(*board)[x][y] = 'B'
	for i := 0; i < 8; i++ {
		px := x + (*dirX)[i]
		py := y + (*dirY)[i]
		if px < 0 || px >= m || py < 0 || py >= n {
			continue
		}
		if (*board)[px][py] == 'E' {
			dfs(px, py, board, m, n, dirX, dirY)
		}
	}
}

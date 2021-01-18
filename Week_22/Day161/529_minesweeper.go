package Day161

func updateBoard(board [][]byte, click []int) [][]byte {
	m, n := len(board), len(board[0])
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	dictX := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dictY := []int{1, 0, -1, 1, -1, 1, 0, -1}
	if board[click[0]][click[1]] != 'E' {
		return board
	}
	queue := [][]int{click}
	for len(queue) > 0 {
		tmp := make([][]int, 0)
		for _, point := range queue {
			x, y := point[0], point[1]
			cnt := 0
			board[x][y] = '.'
			for d := 0; d < 8; d++ {
				nextX, nextY := x+dictX[d], y+dictY[d]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue
				}
				if board[nextX][nextY] == 'M' {
					cnt++
				}
			}
			if cnt > 0 {
				board[x][y] = byte(cnt + '0')
			} else {
				board[x][y] = 'B'
				for d := 0; d < 8; d++ {
					nextX, nextY := x+dictX[d], y+dictY[d]
					if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
						continue
					}
					if board[nextX][nextY] == 'E' {
						tmp = append(tmp, []int{nextX, nextY})
					}
				}
			}
		}
		queue = tmp
	}
	//for i:=0; i<m;i++ {
	//	fmt.Printf("%v\n", string(board[i]))
	//}
	return board
}

package Week_07

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	dictX := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dictY := []int{1, 1, 1, 0, 0, -1, -1, -1}
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	front := make(queue, 0)
	back := make(queue, 0)
	front = append(front, &position{x: 0, y: 0})
	back = append(back, &position{x: n - 1, y: n - 1})
	frontMark, backMark := 2, 3
	grid[0][0] = frontMark
	grid[n-1][n-1] = backMark
	dist := 1
	for len(front) > 0 {
		tmp := make(queue, 0)
		for _, p := range front {
			if grid[p.x][p.y] == backMark {
				return dist
			}
			grid[p.x][p.y] = 1
			for d := 0; d < 8; d++ {
				nextX, nextY := p.x+dictX[d], p.y+dictY[d]
				if nextX < 0 || nextX > n-1 || nextY < 0 || nextY > n-1 || grid[nextX][nextY] == 1 || grid[nextX][nextY] == frontMark {
					continue
				}
				if grid[nextX][nextY] == backMark {
					return dist + 1
				}
				grid[nextX][nextY] = frontMark
				tmp = append(tmp, &position{x: nextX, y: nextY})
			}
		}
		dist++
		front = tmp
		if len(front) > len(back) {
			front, back = back, front
			frontMark, backMark = backMark, frontMark
		}
	}
	return -1
}

type position struct {
	x int
	y int
}

type queue []*position

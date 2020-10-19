package Week_09

func numIslands2(m int, n int, positions [][]int) []int {
	ret := make([]int, len(positions))
	board := make([][]int, m)
	for i := 0; i < m; i++ {
		board[i] = make([]int, n)
	}

	total := m * n
	uf := &unionFind{}
	uf.init(total + 1)
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	for idx, position := range positions {
		x, y := position[0], position[1]
		if board[x][y] == 1 {
			ret[idx] = ret[idx-1]
			continue
		}
		board[x][y] = 1
		//fmt.Printf("uf count %+v\n", uf.count)
		uf.unUnion(x*n + y)
		//fmt.Printf("uf count %+v\n", uf.count)
		for p := 0; p < 4; p++ {
			nextX, nextY := x+dictX[p], y+dictY[p]
			//fmt.Printf("start union (%v)(%v) to (%v)(%v), board(%+v)\n", x, y, nextX, nextY, board[0])
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || board[nextX][nextY] == 0 {
				continue
			}
			//fmt.Printf("end union (%v)(%v) to (%v)(%v)\n", x, y, nextX, nextY)
			uf.union(x*n+y, nextX*n+nextY)
		}
		//fmt.Printf("uf count %+v\n", uf.count)
		ret[idx] = uf.count - 1
	}
	return ret
}

type unionFind struct {
	count   int
	parents []int
}

func (uf *unionFind) init(n int) {
	uf.count = 1
	uf.parents = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = n
	}
}

func (uf *unionFind) find(p int) int {
	tmp := p
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[tmp] != p {
		tmp1 := tmp
		uf.parents[tmp] = p
		tmp = uf.parents[tmp1]
	}
	return p
}

func (uf *unionFind) union(p, q int) {
	p1 := uf.find(p)
	q1 := uf.find(q)
	if p1 == q1 {
		return
	}
	uf.parents[p1] = q1
	uf.count--
}

func (uf *unionFind) unUnion(p int) {
	uf.parents[p] = p
	uf.count++
}

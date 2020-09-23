package Day45

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	uf := &unionFind{}
	total := m * n
	uf.init(total + 1)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				uf.union(i*n+j, total)
			}
			if j < n-1 && board[i][j+1] == 'O' {
				uf.union(i*n+j, i*n+j+1)
			}
			if i < m-1 && board[i+1][j] == 'O' {
				uf.union((i+1)*n+j, i*n+j)
			}
		}
	}
	boss := uf.find(total)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if uf.find(i*n+j) != boss {
				board[i][j] = 'X'
			}
		}
	}
}

type unionFind struct {
	count   int
	parents []int
}

func (uf *unionFind) init(n int) {
	uf.count = n
	uf.parents = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}

func (uf *unionFind) find(p int) int {
	tmp := p
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[tmp] != p {
		tmp1 := tmp
		tmp = uf.parents[tmp]
		uf.parents[tmp1] = p
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

package Day71

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	total := m * n
	uf := &unionFind{}
	uf.init(total + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
				uf.union(i*n+j, total)
			}
			if i < m-1 && board[i][j] == board[i+1][j] {
				uf.union(i*n+j, (i+1)*n+j)
			}
			if j < n-1 && board[i][j] == board[i][j+1] {
				uf.union(i*n+j, i*n+j+1)
			}
		}
	}
	boss := uf.find(total)
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] != 'O' {
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
		uf.parents[tmp] = p
		tmp = uf.parents[tmp1]
	}
	return p
}
func (uf *unionFind) union(p, q int) {
	p1, q1 := uf.find(p), uf.find(q)
	if p1 == q1 {
		return
	}
	uf.parents[p1] = q1
	uf.count--
}

package Week_07

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	uf := new(unionFind130)
	total := m*n + 1
	uf.Init(total)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if board[i][j] == 'O' && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				uf.Union(i*n+j, total-1)
			}
			if j < n-1 && board[i][j+1] == 'O' {
				uf.Union(i*n+j, i*n+j+1)
			}
			if i < m-1 && board[i+1][j] == 'O' {
				uf.Union(i*n+j, (i+1)*n+j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if uf.Find(i*n+j) != uf.Find(total-1) {
				board[i][j] = 'X'
			}
		}
	}

}

type unionFind130 struct {
	count   int
	parents []int
}

func (uf *unionFind130) Init(n int) {
	uf.count = n
	uf.parents = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}
func (uf *unionFind130) Find(p int) int {
	tmp := p
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[tmp] != p {
		parent := uf.parents[tmp]
		uf.parents[tmp] = p
		tmp = parent
	}
	return p
}
func (uf *unionFind130) Union(p, q int) {
	p1 := uf.Find(p)
	q1 := uf.Find(q)
	if p1 == q1 {
		return
	}
	uf.parents[p1] = q1
}

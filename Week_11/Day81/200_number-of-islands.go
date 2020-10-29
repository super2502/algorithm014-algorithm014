package Day81

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	total := m * n
	uf := &unionFind{}
	uf.init(total + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				uf.union(i*n+j, total)
			}
			if i < m-1 && grid[i][j] == grid[i+1][j] {
				uf.union(i*n+j, (i+1)*n+j)
			}
			if j < n-1 && grid[i][j] == grid[i][j+1] {
				uf.union(i*n+j, i*n+j+1)
			}
		}
	}
	return uf.count - 1
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
	for p != uf.parents[p] {
		p = uf.parents[p]
	}
	for uf.parents[tmp] != p {
		tmp1 := uf.parents[tmp]
		uf.parents[tmp] = p
		tmp = tmp1
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

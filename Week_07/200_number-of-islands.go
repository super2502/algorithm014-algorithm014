package Week_07

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	uf := new(unionFind200)
	uf.Init(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				uf.count--
				continue
			}
			if j < n-1 && grid[i][j+1] == '1' {
				uf.Union(i*n+j, i*n+j+1)
			}
			if i < m-1 && grid[i+1][j] == '1' {
				uf.Union(i*n+j, (i+1)*n+j)
			}
		}
	}
	return uf.count
}

type unionFind200 struct {
	count   int
	parents []int
}

func (uf *unionFind200) Init(n int) {
	uf.count = n
	uf.parents = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}

// find时拉平
func (uf *unionFind200) Find(p int) int {
	tmp := p
	for uf.parents[p] != p {
		p = uf.parents[p]
	}
	for uf.parents[tmp] != tmp {
		parent := uf.parents[tmp]
		uf.parents[tmp] = p
		tmp = parent
	}
	return p
}
func (uf *unionFind200) Union(p, q int) {
	p1 := uf.Find(p)
	q1 := uf.Find(q)
	if p1 == q1 {
		return
	}
	uf.parents[p1] = q1
	uf.count--
}

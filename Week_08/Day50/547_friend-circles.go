package Day50

func findCircleNum(M [][]int) int {
	n := len(M)
	uf := new(unionFind)
	uf.init(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.count
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
	p1 := uf.find(p)
	q1 := uf.find(q)
	if p1 == q1 {
		return
	}
	uf.parents[p1] = q1
	uf.count--
}

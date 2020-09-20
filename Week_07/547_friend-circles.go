package Week_07

func findCircleNum(M [][]int) int {
	uf := new(unionFind)
	uf.Init(len(M))
	for i := 0; i < len(M); i++ {
		for j := i; j < len(M); j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count()
}

type unionFind struct {
	count   int
	parents []int
}

func (uf *unionFind) Init(n int) {
	uf.count = n
	uf.parents = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parents[i] = i
	}
}
func (uf *unionFind) Find(p int) int {
	parent := uf.parents[p]
	for parent != p {
		p, parent = parent, uf.parents[parent]
	}
	return p
}
func (uf *unionFind) Union(p, q int) {
	p1 := uf.Find(p)
	q1 := uf.Find(q)
	if p1 == q1 {
		return
	}
	uf.parents[q1] = p1
	for i := 0; i < len(uf.parents); i++ {
		if uf.parents[i] == q1 {
			uf.parents[i] = p1
		}
	}
	uf.count--
}
func (uf *unionFind) Count() int {
	return uf.count
}

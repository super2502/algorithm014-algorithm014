package Day126

import "container/heap"

func trapRainWater(heightMap [][]int) int {
	pq := make(priorityQueue, 0)
	m := len(heightMap)
	if m < 3 {
		return 0
	}
	n := len(heightMap[0])
	if n < 3 {
		return 0
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		heap.Push(&pq, &node{
			height:   heightMap[0][j],
			position: []int{0, j},
		})
		heap.Push(&pq, &node{
			height:   heightMap[m-1][j],
			position: []int{m - 1, j},
		})
		visited[0][j] = true
		visited[m-1][j] = true
	}
	for i := 1; i < m-1; i++ {
		heap.Push(&pq, &node{
			height:   heightMap[i][0],
			position: []int{i, 0},
		})
		heap.Push(&pq, &node{
			height:   heightMap[i][n-1],
			position: []int{i, n - 1},
		})
		visited[i][0] = true
		visited[i][n-1] = true
	}
	dictX := []int{-1, 0, 1, 0}
	dictY := []int{0, 1, 0, -1}
	max := 0
	sum := 0
	for pq.Len() > 0 {
		cn := heap.Pop(&pq).(*node)
		x, y := cn.position[0], cn.position[1]
		if cn.height > max {
			max = cn.height
		}
		for d := 0; d < 4; d++ {
			nextX, nextY := x+dictX[d], y+dictY[d]
			if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || visited[nextX][nextY] {
				continue
			}
			heap.Push(&pq, &node{
				height:   heightMap[nextX][nextY],
				position: []int{nextX, nextY},
			})
			visited[nextX][nextY] = true
			if heightMap[nextX][nextY] < max {
				sum += max - heightMap[nextX][nextY]
			}
		}
	}
	return sum
}

type node struct {
	height   int
	position []int
}
type priorityQueue []*node

func (pq *priorityQueue) Len() int {
	return len(*pq)
}
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].height < (*pq)[j].height
}
func (pq *priorityQueue) Push(x interface{}) {
	n := x.(*node)
	*pq = append(*pq, n)
}
func (pq *priorityQueue) Pop() interface{} {
	n := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return n
}

var _ heap.Interface = &priorityQueue{}

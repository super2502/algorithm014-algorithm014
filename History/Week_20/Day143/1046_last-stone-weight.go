package Day143

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	pq := make(priorityQueue, 0)
	for _, stone := range stones {
		heap.Push(&pq, stone)
	}
	for pq.Len() >= 2 {
		y := heap.Pop(&pq).(int)
		x := heap.Pop(&pq).(int)
		if x < y {
			heap.Push(&pq, y-x)
		}
	}
	if pq.Len() == 1 {
		return pq[0]
	}
	return 0
}

type priorityQueue []int

func (pq *priorityQueue) Len() int {
	return len(*pq)
}
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i] > (*pq)[j]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}

var _ heap.Interface = &priorityQueue{}

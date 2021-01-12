package Day118

import (
	"container/heap"
	"fmt"
)

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	for _, t := range tasks {
		taskMap[t]++
	}
	pq := make(priorityQueue, 0)
	for t, cnt := range taskMap {
		heap.Push(&pq, &task{
			b:   t,
			cnt: cnt,
		})
	}
	//fmt.Printf("%+v\n", pq)
	cost := 0
	for !pq.IsEmpty() {
		//pq.Print()
		k := n + 1
		tmp := make([]*task, 0)
		for !pq.IsEmpty() && k > 0 {
			t := heap.Pop(&pq).(*task)
			tmp = append(tmp, t)
			k--
		}
		for _, t := range tmp {
			t.cnt--
			if t.cnt > 0 {
				heap.Push(&pq, t)
			}
		}
		if pq.IsEmpty() {
			cost += len(tmp)
		} else {
			cost += n + 1
		}
		//fmt.Printf("cost goes to (%v)(%v)\n", cost, n + 1)
	}
	return cost
}

type task struct {
	b   byte
	cnt int
}
type priorityQueue []*task

func (pq *priorityQueue) Len() int {
	return len(*pq)
}
func (pq *priorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].cnt > (*pq)[j].cnt
}
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*task))
}
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}
func (pq *priorityQueue) Print() {
	for _, t := range *pq {
		fmt.Printf(" %v:%v", t.cnt, string(t.b))
	}
	fmt.Printf("\n")
}

package Week_06

import (
	"container/heap"
)
func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	for _, t := range tasks{
		taskMap[t]++
	}
	th := make(taskHeap, 0)
	for t, count := range taskMap {
		heap.Push(&th, &task{
			name: t,
			count: count,
		})
	}
	step := 0
	for th.Len() > 0 {
		round := make([]*task, 0, n)
		for i:=0;i<n+1;i++{
			if th.Len() == 0 {
				break
			}
			t := heap.Pop(&th).(*task)
			t.count--
			if t.count > 0 {
				round = append(round, t)
			}
			step++
		}
		if th.Len() == 0 && len(round) > 0 {
			step += n + 1 - len(round)
		}
		for _, t := range round {
			heap.Push(&th, t)
		}
	}

	return step
}

type task struct {
	name byte
	count int
}

type taskHeap []*task

func (th *taskHeap) Len() int {
	return len(*th)
}

func (th *taskHeap) Less(i, j int) bool {
	return (*th)[i].count > (*th)[j].count
}

func (th *taskHeap) Swap(i, j int) {
	(*th)[i], (*th)[j] = (*th)[j], (*th)[i]
}

func (th *taskHeap) Push(x interface{}) {
	*th = append(*th, x.(*task))
}

func (th *taskHeap) Pop() interface{} {
	x := (*th)[th.Len()-1]
	*th = (*th)[:th.Len()-1]
	return x
}
var _ heap.Interface = &taskHeap{}
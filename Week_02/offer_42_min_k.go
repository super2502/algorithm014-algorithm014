package Week_02

import (
	"container/heap"
	"sort"
)

type thisHeap struct {
	sort.IntSlice
}

func (t *thisHeap) Push(x interface{}) {
	t.IntSlice = append(t.IntSlice, x.(int))
}
func (t *thisHeap) Pop() interface{} {
	if len(t.IntSlice) == 0 {
		return nil
	}
	x := t.IntSlice[len(t.IntSlice)-1]
	t.IntSlice = t.IntSlice[:len(t.IntSlice)-1]
	return x
}

var _ heap.Interface = &thisHeap{}

func MinK(nums []int, k int) []int {

	s := &thisHeap{
		IntSlice: sort.IntSlice(nums),
	}
	heap.Init(s)
	ret := make([]int, 0)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(s).(int))
	}
	return ret
}

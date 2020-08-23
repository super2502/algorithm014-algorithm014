package Week_03

import (
	"container/heap"
	"fmt"
	"sort"
)

type slice struct {
	sort.IntSlice
}

func (s *slice) Push(x interface{}) {
	s.IntSlice = append(s.IntSlice, x.(int))
}

func (s *slice) Pop() interface{} {
	x := s.IntSlice[s.Len()-1]
	s.IntSlice = s.IntSlice[:s.Len()-1]
	return x
}

var _ heap.Interface = &slice{}

func HeapSort() {
	a := []int{3, 4, 6, 1, 2, 5, 8, 1, 0}
	s := &slice{
		IntSlice: sort.IntSlice(a),
	}
	heap.Init(s)

	for s.Len() > 0 {
		fmt.Printf("%v\n", heap.Pop(s))
	}
}

func HeapSortNode() {
	s := make(MyHeap, 0)
	heap.Init(&s)
	heap.Push(&s, &heapNode{
		num:   555,
		count: 1,
	})
	heap.Push(&s, &heapNode{
		num:   666,
		count: 2,
	})
	heap.Push(&s, &heapNode{
		num:   777,
		count: 3,
	})
	heap.Push(&s, &heapNode{
		num:   888,
		count: 2,
	})

	for s.Len() > 0 {
		fmt.Printf("%+v\n", heap.Pop(&s))
	}
}

type heapNode struct {
	num   int
	count int
}
type MyHeap []*heapNode

func (mh *MyHeap) Len() int {
	return len(*mh)
}

// 大顶堆
func (mh *MyHeap) Less(i, j int) bool {
	return (*mh)[i].count > (*mh)[j].count
}
func (mh *MyHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}
func (mh *MyHeap) Push(x interface{}) {
	*mh = append(*mh, x.(*heapNode))
}
func (mh *MyHeap) Pop() interface{} {
	x := (*mh)[mh.Len()-1]
	*mh = (*mh)[:mh.Len()-1]
	return x
}

var _ sort.Interface = &MyHeap{}
var _ heap.Interface = &MyHeap{}

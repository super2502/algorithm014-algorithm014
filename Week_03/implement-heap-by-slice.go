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

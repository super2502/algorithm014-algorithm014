package Week_02

import (
	"container/heap"
	"fmt"
)

var _ = fmt.Errorf("")

type node struct {
	val   int
	count int
}

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	a := &A{}
	k = min(len(m), k)
	if k == 0 {
		return []int{}
	}
	fmt.Printf("nums (%+v)\n", nums)
	fmt.Printf("m (%+v)\n", m)
	j := 0
	for val, count := range m {
		fmt.Printf("val, count (%v)(%v), j (%v) k (%v)\n", val, count, j, k)
		fmt.Printf("a (%+v) , len(a)(%v)\n", a, a.Len())
		n := &node{
			val:   val,
			count: count,
		}
		if j < k {
			heap.Push(a, n)
			j++
			continue
		}
		if count > (*a)[0].count {
			heap.Pop(a)
			heap.Push(a, n)
		}
	}

	ret := make([]int, 0)
	for i := 0; i < k; i++ {

		ret = append(ret, (*a)[i].val)

	}
	return ret
}

type A []*node

func (a A) Less(i, j int) bool {
	return a[i].count < a[j].count
}

func (a A) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]

}

func (a A) Len() int {
	return len(a)
}

func (a *A) Push(x interface{}) {

	*a = append(*a, x.(*node))
}

func (a *A) Pop() interface{} {
	l := a.Len()
	x := (*a)[l-1]
	*a = (*a)[:l-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

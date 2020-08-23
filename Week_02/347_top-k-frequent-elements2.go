package Week_02

import (
	"container/heap"
)

// 分别用大顶堆，小顶堆计算
func topKFrequent2(nums []int, k int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 || k == 0 {
		return ret
	}
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	mh := make(MyHeap, 0)
	heap.Init(&mh)
	for num, count := range counter {
		if mh.Len() == k {
			if mh[0].count < count {
				heap.Pop(&mh)
				heap.Push(&mh, &heapNode{
					num:   num,
					count: count,
				})
			}
		} else {
			heap.Push(&mh, &heapNode{
				num:   num,
				count: count,
			})
		}

	}
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(&mh).(*heapNode).num)
	}
	return ret
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
// func (mh *MyHeap) Less(i, j int) bool {
// 	return (*mh)[i].count > (*mh)[j].count
// }
// 小顶堆
func (mh *MyHeap) Less(i, j int) bool {
	return (*mh)[i].count < (*mh)[j].count
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

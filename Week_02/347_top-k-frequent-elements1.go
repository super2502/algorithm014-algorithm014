package Week_02

import (
	"container/heap"
	"sort"
)

func topKFrequent1(nums []int, k int) []int {

	freq := make(map[int]int)
	freqMap := make(map[int][]int)

	for _, num := range nums {
		freq[num]++
	}
	counts := make([]int, 0, len(freq))
	for num, count := range freq {
		if _, ok := freqMap[count]; !ok {
			freqMap[count] = make([]int, 0)
		}
		freqMap[count] = append(freqMap[count], num)
		counts = append(counts, count)
	}

	mh := &myHeap{
		IntSlice: &sort.IntSlice{},
	}
	for i := 0; i < k; i++ {
		heap.Push(mh, counts[i])
	}
	for i := k; i < len(counts); i++ {
		if counts[i] > (*mh.IntSlice)[0] {
			heap.Pop(mh)
			heap.Push(mh, counts[i])
		}
	}
	ret := make([]int, 0, k)
	for i := 0; i < k && len(ret) < k; i++ {
		ret = append(ret, freqMap[heap.Pop(mh).(int)]...)
	}
	return ret
}

type myHeap struct {
	*sort.IntSlice
}

var _ heap.Interface = &myHeap{}

func (h *myHeap) Push(x interface{}) {
	*h.IntSlice = append(*h.IntSlice, x.(int))
}

func (h *myHeap) Pop() interface{} {
	x := (*h.IntSlice)[len(*h.IntSlice)-1]
	*h.IntSlice = (*h.IntSlice)[:len(*h.IntSlice)-1]
	return x
}

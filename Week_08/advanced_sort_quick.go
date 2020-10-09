package Week_08

import (
	//"fmt"
)
func quickSort(nums []int, i, j int) {

	//fmt.Printf("in loop (%v)(%v)\n", i, j)
	if i >= j {
		return
	}
	pivot := partition(nums, i, j)
	//fmt.Printf("get pivot (%v)(%v)(%v) (%+v)\n", i, j, pivot, nums)
	quickSort(nums,i ,pivot-1)
	quickSort(nums, pivot+1, j)

}

func partition(nums []int, i, j int) int{

	pivot := j
	start := i
	for k := i; k < j; k++ {
		if nums[k] < nums[j] {
			nums[start], nums[k] = nums[k], nums[start]
			start++
		}
	}
	nums[pivot], nums[start] = nums[start], nums[pivot]
	return start
}

func quickSortIteration(nums []int) {

	s := make(stack, 0)
	s.Push([]int{0, len(nums)-1})
	for !s.IsEmpty() {
		x := s.Pop().([]int)
		start, end := x[0], x[1]
		if start >= end {
			continue
		}
		pivot := end
		idx := start
		for i := start; i < end; i++ {
			if nums[i] < nums[pivot] {
				nums[idx], nums[i] = nums[i], nums[idx]
				idx++
			}
		}
		nums[pivot], nums[idx] = nums[idx], nums[pivot]
		s.Push([]int{start, idx - 1})
		s.Push([]int{idx + 1, end})
	}
}

type stack []interface{}

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0
}
func (s *stack) Push(x interface{}) {
	*s = append(*s, x)
}
func (s *stack) Pop() interface{} {
	x := (*s)[s.Len() - 1]
	*s = (*s)[:s.Len() - 1]
	return x
}
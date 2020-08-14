package Week_01

import (
	"fmt"
)

func trap(height []int) int {

	fmt.Printf("%+v\n", height)
	s := &stack{}
	l := len(height)
	if l <= 2 {
		return 0
	}
	water := 0

	for i := 0; i < l; i++ {
		for !s.isEmpty() && height[s.top()] < height[i] {
			top := s.pop()
			if s.isEmpty() {
				break
			}
			h := min(height[s.top()], height[i])
			water += (h - height[top]) * (i - s.top() - 1)
		}
		//fmt.Printf("push i (%v)\n", i)
		s.push(i)
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop() int {
	n := []int(*s)[len(*s)-1]
	*s = []int(*s)[:len(*s)-1]
	return n
}

func (s *stack) top() int {
	return []int(*s)[len(*s)-1]
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

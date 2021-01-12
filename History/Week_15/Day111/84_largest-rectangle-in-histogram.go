package Day111

import "fmt"

func largestRectangleArea(heights []int) int {
	s := make(stack, 0)
	s.push(-1)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		if !s.isEmpty() {
			fmt.Printf("i (%v)(%v), (%+v) (%v)\n", i, heights[i], s, heights[s.peek()])
		}
		for !s.isEmpty() && heights[i] <= heights[s.peek()] {
			idx := s.pop()
			maxArea = max(maxArea, heights[idx]*(i-s.peek()-1))
		}
		s.push(i)
	}
	for !s.isEmpty() {
		idx := s.pop()
		maxArea = max(maxArea, heights[idx]*(len(heights)-s.peek()-1))
	}
	return maxArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type stack []int

func (s *stack) len() int {
	return len(*s)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0 || (*s)[s.len()-1] == -1
}
func (s *stack) push(x int) {
	*s = append(*s, x)
}
func (s *stack) pop() int {
	x := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return x
}
func (s *stack) peek() int {
	x := (*s)[s.len()-1]
	return x
}

package main

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	s := make(stack, 0)
	area := 0
	s.Push(-1)
	for i := 0; i < len(heights); i++ {
		for !s.IsEmpty() && heights[i] <= heights[s.Peek()] {
			idx := s.Pop()
			area = max(area, heights[idx]*(i-s.Peek()-1))
		}
		s.Push(i)
	}
	for !s.IsEmpty() {
		idx := s.Pop()
		area = max(area, heights[idx]*(len(heights)-s.Peek()-1))
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {

	if len(height) == 0 {
		return 0
	}
	s := make(stack, 0)
	sum := 0
	for i := 0; i < len(height); i++ {
		for !s.IsEmpty() && height[i] > height[s.Peek()] {
			idx := s.Pop()
			if !s.IsEmpty() {
				sum += (min(height[i], height[s.Peek()]) - height[idx]) * (i - s.Peek() - 1)
			}
		}
		s.Push(i)
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type stack []int

func (s *stack) Len() int {
	return len(*s)
}
func (s *stack) IsEmpty() bool {
	return s.Len() == 0 || (*s)[s.Len()-1] == -1
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Pop() int {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack) Peek() int {
	x := (*s)[s.Len()-1]
	return x
}

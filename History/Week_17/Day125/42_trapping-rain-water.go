package Day125

func trap(height []int) int {
	s := make(stack, 0)
	area := 0
	for i := 0; i < len(height); i++ {
		for !s.IsEmpty() && height[i] >= height[s.Peek()] {
			idx := s.Pop()
			if !s.IsEmpty() {
				area += (i - s.Peek() - 1) * (min(height[i], height[s.Peek()]) - height[idx])
			}
		}
		s.Push(i)
	}
	return area
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
	return s.Len() == 0
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

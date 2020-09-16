package Week_05

func longestValidParentheses(s string) int {
	st := make(stack, 0, len(s))
	st.Push(-1)
	maxLength := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st.Push(i)
		} else {
			st.Pop()
			if !st.IsEmpty() {
				maxLength = max32(maxLength, i-st.Seek())
			} else {
				st.Push(i)
			}
		}
	}
	return maxLength

}

func max32(a, b int) int {
	if a > b {
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
func (s *stack) Pop() int {
	x := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return x
}
func (s *stack) Push(x int) {
	*s = append(*s, x)
}
func (s *stack) Seek() int {
	return (*s)[s.Len()-1]
}
